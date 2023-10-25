package behavior

// declaration la plus générique possible pour un element
// susceptible d'être visité
// impossible de faire mieux
// j'ai pris la convention C++ pour la declaration des interfaces ici
// je sais c'est lourd et old school
// IVisitable, IProduct, IVisitor
// c'est juste pour l'exercice
// par contre la convention en go c'est ajouter le suffixe er
// https://go.dev/doc/effective_go#interface-names
// la norme go serait
// Visitabler (?), Producter (??, contre-sens ici) , Visitorer (aie , horreur)
// pour la suite je garderai Visitable, Product, Visitor  (ni -er quand c'est pas possible , ni I partout)
// j'ai encore un fort accent C++...

type IVisitable interface {
	Accept(IVisitor)
}

// l'interface d'un produit
type IProduct interface {
	GetName() string
	GetPrice() float32
}

type IVisitor interface {
	Visit(IProduct)
}

// je n'ai pas de classe au sens de l'objet et je ne m'en porte que mieux !!!
// après des decennies de pratiques objet
// I don't care about OOP, DON'T NEED THAT IN SERVICE ARCHITECTURE SOFTWARE
// See :https://confluence.oodrive.net/display/ARCH/Styles+d%27Architecture
// Je n'engage plus de conversation désormais avec les gens qui me parlent
// de programmation objet pour les services backend d'un SI

type Product struct {
	Price float32
	Name  string
}

// Product va implementer les 2 interfaces IVisitable et IProduct

func (p *Product) GetPrice() float32 {
	return p.Price
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) Accept(v IVisitor) {
	v.Visit(p)
}

type OmegaSpeedMaster struct {
	Product
}

type OmegaSeaMaster struct {
	Product
}

// mon visiteur qui calcule les prix

type PriceVisitor struct {
	Sum float32
}

func (pv *PriceVisitor) Visit(p IProduct) {
	pv.Sum += p.GetPrice()
}
