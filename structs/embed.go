package main


type contactInfo struct {
	email string
	zipCode int
}

type personC struct {
	firstName string
	lastName string
	contact contactInfo
}


func returnpersonC() personC{
	jim := personC{
		firstName: "Jim",
		lastName: "Halpert",
		contact: contactInfo{
			email: "jimhalpert@dmf.com",
			zipCode: 20010,
		},
	}

	return jim
}
