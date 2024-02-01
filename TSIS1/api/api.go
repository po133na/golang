package api

type Breed struct {
	Id        int    `json:"id"`
	BreedName string `json:"breed_name"`
	Country   string `json:"country"`
}

var Breeds = []Breed{

	{1, "Persian", "Iran"},
	{2, "Siamese", "Thailand"},
	{3, "Maine Coon", "United States"},
	{4, "Bengal", "United States"},
	{5, "Sphynx", "Canada"},
	{6, "Ragdoll", "United States"},
	{7, "British Shorthair", "United Kingdom"},
	{8, "Abyssinian", "Ethiopia"},
	{9, "Scottish Fold", "United Kingdom"},
	{10, "Russian Blue", "Russia"},
	{11, "Turkish Angora", "Turkey"},
	{12, "Norwegian Forest", "Norway"},
	{13, "Burmese", "Burma"},
	{14, "Oriental Shorthair", "United States"},
	{15, "Devon Rex", "United Kingdom"},
	{16, "Himalayan", "United States"},
	{17, "Manx", "Isle of Man"},
	{18, "Cornish Rex", "United Kingdom"},
	{19, "Chartreux", "France"},
	{20, "Japanese Bobtail", "Japan"},
}
