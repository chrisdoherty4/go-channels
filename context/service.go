
func httpget(request Request) Result



func dispatchRequests(requests []*Requests) []*Result {
	results := make([]*Result, len(requests))

	for request := range requests {
		result := httpget(request)
		results = append(results, resullt)
	}

	return results
}


func httpget(request Request) <-chan Result

func merge(channels []chan Result) <-chan Result

func dispatchRequests(requests []*Requests) []*Result {
	resultChannels := make([]chan Result)

	for request := range requests {
		resultChannels = append(resultsChannels, httpget(request))
	}

	results := make([]*Result, len(requests))

	for result := merge(resultChannels) {
		results = append(results, result)
	}

	return results
}



func httpget(ctx context.Context, request Request) <-chan Result

func merge(channels []chan Result) <-chan Result

func dispatchRequests(ctx context.Context, requests []*Requests) ([]*Result, error) {
	resultChannels := make([]chan Result)

	for request := range requests {
		resultChannels = append(resultsChannels, httpget(ctx, request))
	}

	results := make([]*Result, len(requests))

	for {
		select {
		case ctx.Done():
			return nil, ctx.Err()
		case result := merge(resultChannels):
			results = append(results, result)
		}
	}


	return results, nil
}
