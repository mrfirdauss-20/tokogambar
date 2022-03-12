package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/corona10/goimagehash"
	"github.com/riandyrn/go-knn"
)
import _ "image/jpeg"

const addr = ":7124"

func main() {
	// initialize database
	dbRecords, err := loadDB()
	if err != nil {
		log.Fatalf("unable to initialize database due: %v", err)
	}
	// attach handler
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./images"))))
	http.HandleFunc("/similars", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			WriteAPIResp(w, NewErrorResp(NewErrNotFound()))
			return
		}
		// parse request body
		var rb searchReqBody
		err := json.NewDecoder(r.Body).Decode(&rb)
		if err != nil {
			WriteAPIResp(w, NewErrorResp(NewErrBadRequest(err.Error())))
			return
		}
		// validate request body
		err = rb.Validate()
		if err != nil {
			WriteAPIResp(w, NewErrorResp(err))
			return
		}
		// search similar images
		imgData, err := rb.GetByte()
		if err != nil {
			WriteAPIResp(w, NewErrorResp(err))
			return
		}
		similarImages, err := searchSimilarImages(dbRecords, imgData)
		if err != nil {
			WriteAPIResp(w, NewErrorResp(err))
			return
		}
		// output success response
		WriteAPIResp(w, NewSuccessResp(similarImages))
	})
	// start server
	log.Printf("server is listening on %v", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("unable to start server due: %v", err)
	}
}

func loadDB() (*(knn.KNN), error) {
	basePath := "./images"
	infos, err := ioutil.ReadDir(basePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read dir due: %w", err)
	}
	var filenames []string
	for _, info := range infos {
		if info.IsDir() {
			continue
		}
		filenames = append(filenames, info.Name())
	}
	knn := knn.NewKNN(knn.Configs{
		VectorDimension: 18,
		NumHashTable:    3,
		NumHyperplane:   3,
		SlotSize:        5,
	})
	docs := []ImageDoc{}
	for _, filename := range filenames {
		b, err := ioutil.ReadFile(basePath + "/" + filename)
		if err != nil {
			continue
		}
		hash,_:=goimagehash.PerceptionHash(byteToImg(b))
		arr :=  srtringBinToArrFloat(hash.ToString())
		docs=  append(docs,ImageDoc{
			ID: filename,
			Vector: arr,
		})
	}
	for i:=0;i<len(docs);i++{
		err:=knn.Add(&docs[i])
		if err != nil{
			log.Fatalf("unable to add new document to knn index due: %v",err)
		}
	}
	
	return knn, nil
}

func searchSimilarImages(knn *(knn.KNN), data []byte) ([]similarImage, error) {
	//hashStr := getHash(data)	
	//imag, _, _ := image.Decode(bytes.NewReader(data))
	baseHash,_ := goimagehash.PerceptionHash(byteToImg(data))
	vec := srtringBinToArrFloat(baseHash.ToString())
	resultDocs, err:=knn.Query(vec,1)
	if err!=nil{
		log.Printf("unable to query documents due:%v",err)
	}
	simImages := []similarImage{}
	for _, resultDoc := range resultDocs {
		simImages = append(simImages, similarImage{
				FileName:        resultDoc.Document.GetID(),
				SimilarityScore:  resultDoc.Distance,
		})
	}
	return simImages, nil
}


func byteToImg (data []byte) image.Image{
	img, err := jpeg.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatalln(err)
		fmt.Println("di sini")
	}	
	return img
}

func srtringBinToArrFloat(str string) []float64{
	var ret[]float64
	for _, run := range str{
		ret=append(ret,float64(run)-float64('0'))
	}

	return ret
}