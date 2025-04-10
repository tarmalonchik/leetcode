package main

func numUniqueEmails(emails []string) int {
	mp := make(map[string]interface{})
	for i := range emails {
		mp[toSameFormat(emails[i])] = nil
	}
	return len(mp)
}

func toSameFormat(in string) string {
	out := make([]byte, 0)

	ignoreTillAt := false
	domainPartStarted := false
	for i := range in {
		if in[i] == '@' {
			domainPartStarted = true
			out = append(out, in[i])
		}

		if ignoreTillAt {
			if in[i] == '@' {
				out = append(out, in[i])
				ignoreTillAt = false
				continue
			} else {
				continue
			}
		}

		if in[i] == '.' && !domainPartStarted {
			continue
		}

		if in[i] == '+' && !domainPartStarted {
			ignoreTillAt = true
			continue
		}
		out = append(out, in[i])
	}
	return string(out)
}
