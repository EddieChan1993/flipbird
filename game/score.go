package game

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
)

const storagePath = "./data.json"
const Salt = "DC"

type IScoreDb interface {
	Save(score int)
	GetBestScore() int
	GetLastScore() int
}

type ScoreDb struct {
	BestScore int
	NowScore  int
	KeyHash   string
}

var OSFile *os.File

func init() {
	f, err := os.OpenFile(storagePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	OSFile = f
}

func Init() IScoreDb {
	s := &ScoreDb{}
	fd, err := ioutil.ReadAll(OSFile)
	if err != nil {
		log.Fatal(err)
	}
	if len(fd) == 0 {
		return s
	}
	err = json.Unmarshal(fd, s)
	if err != nil {
		log.Fatal(err)
	}
	if s.md5(s.BestScore) != s.KeyHash {
		log.Fatal("game score stolen！")
	}
	return s
}

func (s *ScoreDb) Save(score int) {
	s.BestScore = int(math.Max(float64(score), float64(s.BestScore)))
	s.NowScore = score
	s.KeyHash = s.md5(s.BestScore)
	contentByte, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	OSFile.Truncate(0)
	n, _ := OSFile.Seek(0, 0)
	_, err = OSFile.WriteAt(contentByte, n)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *ScoreDb) GetBestScore() int {
	return s.BestScore
}

func (s *ScoreDb) GetLastScore() int {
	return s.NowScore
}

func (s *ScoreDb) md5(score int) string {
	h := md5.New()
	h.Write([]byte(strconv.Itoa(score) + Salt))
	return hex.EncodeToString(h.Sum(nil))
}
