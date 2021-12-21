package main

import (
	"fmt"
	"server-mk3/internal/services/idea_request"
)

func main() {
	ideaRequestService, err := idea_request.NewIdeaRequestService(
		idea_request.WithMemoryIdeaRequestRepository(),
	)
	if err != nil {
		panic(err)
	}
	ir, err := ideaRequestService.CreateIdeaRequest("Test the idea", "Wow")
	if err != nil {
		panic(err)
	}
	fmt.Println(ir.GetID(), ": ", ir.GetUpvotes())

	irNew, err := ideaRequestService.UpvoteIdeaRequest(ir.GetID())
	if err != nil {
		panic(err)
	}
	fmt.Println(irNew.GetID(), ": ", irNew.GetUpvotes())
	list, _ := ideaRequestService.ListAllIdeas()
	for _, i := range list {
		fmt.Println(i.GetIdea().Name, ": ", i.GetUpvotes())
	}
	// pass service to handler/resolver

	// start server
}
