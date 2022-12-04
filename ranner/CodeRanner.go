package ranner

import (
	"My-Exercise/model/dto"
	"bytes"
	"io"
	"log"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

func RannerCode(codeDTO *dto.CodeDTO, passCount *int, mu *sync.Mutex, answerError chan struct{}, compileError chan struct{}, maxMemError chan struct{}) {
	cmd := exec.Command("go", "run", codeDTO.Path)
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	pipe, err := cmd.StdinPipe()
	if err != nil {
		return
	}

	_, err = io.WriteString(pipe, codeDTO.Input)
	if err != nil {
		return
	}

	var bm runtime.MemStats
	runtime.ReadMemStats(&bm)
	start := time.Now()
	err = cmd.Run()
	log.Printf("code rannner cost %v", time.Since(start).Milliseconds())
	var em runtime.MemStats
	runtime.ReadMemStats(&em)
	if err != nil {
		log.Printf("code ranner compile error: %v", err)
		compileError <- struct{}{}
		return
	}
	useMem := (em.Alloc - bm.Alloc) / 1024
	if useMem > uint64(codeDTO.MaxMem) {
		log.Printf("code ranner maxMem error, useMem: %v", useMem)
		maxMemError <- struct{}{}
		return
	}

	if out.String() != codeDTO.Output {
		log.Printf("code ranner answer error, answer: %v", out.String())
		answerError <- struct{}{}
		return
	}

	mu.Lock()
	defer mu.Unlock()
	*passCount = *passCount + 1
	return
}
