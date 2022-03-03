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

func loadDB() ([]dbRecord, error) {
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
	var dbRecords []dbRecord
	for _, filename := range filenames {
		b, err := ioutil.ReadFile(basePath + "/" + filename)
		if err != nil {
			continue
		}

		dbRecords = append(dbRecords, dbRecord{
			FileName: filename,
			Ima:	  b,
		})
	}
	return dbRecords, nil
}

func searchSimilarImages(dbRecords []dbRecord, data []byte) ([]similarImage, error) {
	//hashStr := getHash(data)
	//imag, _, _ := image.Decode(bytes.NewReader(data))
	baseHash,_ := goimagehash.PerceptionHash(byteToImg(data))
	simImages := []similarImage{}
	
	for _, record := range dbRecords {
		imgHash,_ := goimagehash.PerceptionHash(byteToImg(record.Ima))
		distance,_:=baseHash.Distance(imgHash)
		if distance<5{
			simImages = append(simImages, similarImage{
				FileName:        record.FileName,
				SimilarityScore: (float64) (distance),
			})
			//mt.Println("distance %d",distance)
		}
	}
	return simImages, nil
}

/*
func getHash(data []byte) ImageHash {
	ima,_,_:=image.Decode(bytes.NewReader(data))
	has,_:=goimagehash.PerceptionHash(ima)
	h:=has.GetHash()
	k:=has.GetKind()
	imageHash:= ImageHash{
		hash: h,
		kind: k,
	}
	return imageHash
}
*/

func byteToImg (data []byte) image.Image{
	img, err := jpeg.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatalln(err)
		fmt.Println("di sini")
	}	
	return img
}