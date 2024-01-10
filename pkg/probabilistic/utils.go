package probabilistic

import "github.com/spaolacci/murmur3"

var SimpleInput = []string{
	"the,",
	"quick",
	"fox,",
	"jumps",
	"over",
	"the",
	"lazy",
	"dog",
	"\n",
	"random",
	"strings",
	"lol",
}

var Tests = []struct {
	Input    string
	Expected bool
}{
	{
		Input:    "hello",
		Expected: false,
	},
	{
		Input:    "world",
		Expected: false,
	}, {
		Input:    "good",
		Expected: false,
	}, {
		Input:    "morning",
		Expected: false,
	},
	{
		Input:    "fox",
		Expected: false,
	},
	{
		Input:    "the,",
		Expected: true,
	},
	{
		Input:    "quick",
		Expected: true,
	},
	{
		Input:    "quickly",
		Expected: true,
	},
	{
		Input:    "fox,",
		Expected: true,
	}, {
		Input:    "jumps",
		Expected: true,
	}, {
		Input:    "over",
		Expected: true,
	}, {
		Input:    "the",
		Expected: true,
	}, {
		Input:    "lazy",
		Expected: true,
	}, {
		Input:    "dog",
		Expected: true,
	},
}

var Words = []string{
	"brake",
	"hello",
	"world",
	"outlaw",
	"intercede",
	"diplomacy",
	"kick",
	"cripple",
	"those",
	"rob",
	"legislative",
	"unacceptable",
	"eloquent",
	"telephone",
	"imaging",
	"uphill",
	"fund",
	"pervade",
	"incidence",
	"sensuality",
	"inconsequential",
	"voter",
	"pubic",
	"skyrocket",
	"entitlement",
	"midst",
	"fat",
	"special",
	"rejection",
	"pare",
	"perversion",
	"briefcase",
	"norm",
	"present",
	"perfect",
	"unrecognized",
	"propensity",
	"litter",
	"illustrate",
	"improvisation",
	"eminently",
	"accountable",
	"dosage",
	"abstract",
	"adjusted",
	"unruly",
	"pivot",
	"rotating",
	"etiology",
	"nearby",
	"bus station",
	"drink",
	"pitch",
	"sclerosis",
	"progress",
	"uh-huh",
	"underserved",
	"slug",
	"caring",
	"sometime",
	"sake",
	"fix",
	"clot",
	"humanism",
	"checkout",
	"give",
	"lettuce",
	"proudly",
	"grass",
	"shiloh",
	"jaliyah",
	"toothless",
	"harvest",
	"rat",
	"drug",
	"buttress",
	"contagion",
	"disappoint",
	"cartilage",
	"further",
	"part-time",
	"scout",
	"rickety",
	"summons",
	"gardening",
	"maritime",
	"connector",
	"mckinley",
	"stay",
	"presence",
	"relevant",
	"situate",
	"consummate",
	"oppressed",
	"originate",
	"propel",
	"oubliette",
	"jeep",
	"glass",
	"legality",
	"increased",
	"brake",
}

func max(x uint, y uint) uint {
	if x > y {
		return x
	}
	return y
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func baseHashes(data []byte) [4]uint64 {
	a1 := []byte{1}
	hasher := murmur3.New128()
	_, err := hasher.Write(data)
	check(err)

	v1, v2 := hasher.Sum128()
	_, err = hasher.Write(a1)
	check(err)

	v3, v4 := hasher.Sum128()
	return [4]uint64{
		v1, v2, v3, v4,
	}
}
