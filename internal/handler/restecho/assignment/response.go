package restecho

type AssignmentResponse struct {
	Title        string
	Point        int
	Intruction   string
	Refference   string
	MyAssignment string
}

func getAssignment(res AssignmentResponse) AssignmentResponse {
	return AssignmentResponse{
		Title:        res.Title,
		Point:        res.Point, //sementara
		Intruction:   res.Intruction,
		Refference:   res.Refference,
		MyAssignment: res.MyAssignment,
	}
}
