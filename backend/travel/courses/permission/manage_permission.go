package permission

type permission struct {
	roles   []string
	methods []string
}

type authority map[string]permission

var authorities = authority{
	"/user": permission{
		roles:		[]string{"USER"},
		methods:	[]string{"GET", "PUT", "DELETE"},
	},
	"/api": permission{
		roles:		[]string{"USER"},
		methods:	[]string{"GET"},
	},
	"/admin": permission{
		roles:		[]string{"ADMIN"},
		methods:	[]string{"GET", "POST", "PUT", "DELETE"},
	},
	"/login": permission{
		roles:   []string{"USER", "ADMIN"},
		methods: []string{"POST"},
	},
	"/logout": permission{
		roles:   []string{"USER", "ADMIN"},
		methods: []string{"POST"},
	},
	"/signup": permission{
		roles:   []string{"USER"},
		methods: []string{"POST"},
	},
}