package main

import (
	"geekbang_study/proto"
	"io"
	"log"
	"strconv"
)

type StreamService struct {
}

func (s *StreamService) Conversations(srv proto.Stream_ConversationsServer) error {
	n := 1
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = srv.Send(&proto.StreamResponse{
			Answer: "from stream server answer: the " + strconv.Itoa(n) + " question is " + req.Question,
		})
		if err != nil {
			return err
		}
		n++
		log.Printf("from stream client question: %s\n", req.Question)
	}
}
