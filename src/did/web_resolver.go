package did

// import (
// 	. "github.com/ockam-network/did"
// )

type WebResolver struct{}

func (w WebResolver) Resolve(did string) ([]byte, error) {
	return nil, nil
}

// req, err := http.NewRequest(http.MethodGet, url, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	resBody, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return resBody
