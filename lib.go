package librsyncdartwrapper

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"strconv"

	"github.com/balena-os/librsync-go"
	log "github.com/sirupsen/logrus"
)

type CallbackWriter struct {
	cb func([]byte, int)
}

func (w *CallbackWriter) Write(b []byte) (int, error) {
	w.cb(b, len(b))
	return len(b), nil
}

func ComputeDelta(signatureStr string, targetFilePath string, writeCallback func([]byte, int)) {
	var m map[string]interface{}
	var sig *librsync.SignatureType
	if signatureStr != "" {
		err := json.Unmarshal([]byte(signatureStr), &m)
		if err != nil {
			log.Fatalf("[librsync]: Failed to parse signature: %v", err)
			return
		}
		strongSigsBase64 := m["StrongSigs"].([]interface{})
		strongSigs := make([][]byte, len(strongSigsBase64))
		for idx, str := range strongSigsBase64 {
			b, err := base64.StdEncoding.DecodeString(str.(string))
			if err != nil {
				log.Fatalf("[librsync]: Failed to decode base64 strongSig: %v", err)
			}
			strongSigs[idx] = b
		}

		weak2BlockRaw := m["Weak2block"].(map[string]interface{})
		weak2Block := make(map[uint32]int)
		for k, v := range weak2BlockRaw {
			i, err := strconv.ParseInt(k, 10, 32)
			if err != nil {
				log.Fatalf("[librsync]: Failed to convert weak2Block key from string -> int: %v", err)
			}
			weak2Block[uint32(i)] = int(v.(float64))
		}

		sig = &librsync.SignatureType{
			SigType:    librsync.MagicNumber(m["SigType"].(float64)),
			BlockLen:   uint32(m["BlockLen"].(float64)),
			StrongLen:  uint32(m["StrongLen"].(float64)),
			StrongSigs: strongSigs,
			Weak2block: weak2Block,
		}
	}

	file, err := os.OpenFile(targetFilePath, os.O_RDONLY, 04444)
	if err != nil {
		log.Fatalf("[librsync]: Failed to open file '%v': %v", targetFilePath, err)
		return
	}
	defer file.Close()

	callbackWriter := &CallbackWriter{
		cb: writeCallback,
	}
	log.Debugf("WriteCallback: %v", writeCallback)
	librsync.Delta(sig, file, callbackWriter)
	log.Debugf("[librsync]: Finished computing delta")
}
