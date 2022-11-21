package planet
import (
	"errors"
	"fmt"
	"log"
	"strings"
	"alieninvasion/src/github/line"
)
const (
	MaxOccupancy = 2
	//Max occupancy of aliens that may occupy a city
	MaxEdges = 4

)
type Map struct {
	cities map[string] *City 
	aliens map[string] *Alien
}
type City struct {
	name string 
	inLinks map[string]string 
	outLinks map[string]string 
	alienOccupancy map[string]*Alien 
}
func (c *City) Priority(other interface{}) bool {
	if t, ok := other.(*City); ok {
		return len(c.outLinks) > len(t.outLinks)
	}
	return false 
}
func NewMap() *Map {
	return &Map{
		cities: make(map[string]*City),
		aliens: make(map[string]*Alien),
	}
}
func (m *Map) AlienNames() []string  {
	alienNames := make([]string, 0, len(m.aliens))
	for alienName:= range m.aliens {
		alienNames = append(alienNames,  alienName)
	}
	return alienNames
}
func (m *Map) Cities() []*City {
	cities := make([]*City, 0, len(m.cities))
	for _, city := range m.cities {
		cities = append(cities, city)
	}
	return cities 
}
func (m *Map) NumCities() uint {
	return uint(len(m.cities))
}

func (m *Map) NumAliens() uint {
	return uint(len(m.aliens))
}
func (m *Map) CityNames() []string {
	cityNames := make([]string, 0, m.NumCities())
	for cityName := range m.cities {
		cityNames = append(cityNames, cityName)
	}
	return cityNames
}
func (m *Map) AddLink(cityName, linkCityDir, linkCityName string) {
	if _, ok:= m.cities[cityName]; !ok {
		m.cities[cityName] = &City{
			name: linkCityName,
			inLinks: make(map[string]string, MaxEdges),
			outLinks: make(map[string]string, MaxEdges),
			alienOccupancy: make(map[string]*Alien, MaxOccupancy),
		}
	}
	if _, ok:= m.cities[linkCityName]; !ok {
		m.cities[linkCityName] = &City{
			name: linkCityName,
			inLinks: make(map[string]string, MaxEdges),
			outLinks: make(map[string]string, MaxEdges),
			alienOccupancy: make(map[string]*Alien, MaxOccupancy),
		}
	}
	m.cities[cityName].outLinks[strings.ToLower(linkCityDir)] = linkCityName
	m.cities[linkCityName].inLinks[strings.ToLower(linkCityDir)] = cityName
}
func(m *Map) MoveAlien() (string, error) {
	for _, alien := range m.aliens {
		occupiedCity := alien.city 
		city := m.cities[occupiedCity]
		for _, linkCityName := range city.outLinks {
			linkCity := m.cities[linkCityName]
			if len(linkCity.alienOccupancy) < MaxOccupancy {
				delete(city.alienOccupancy, alien.name)
				alien.city = linkCity.name
				linkCity.alienOccupancy[alien.name] = alien 
				return alien.name, nil 
			}
		}
	}
	return "", errors.New("CANNOT MOVE")
}
func (m *Map) destroyCity(city *City) []string {
	destroyedAliens := make([]string, 0, MaxOccupancy)
	for alienName := range city.alienOccupancy{
		destroyedAliens = append(destroyedAliens, alienName)
		delete(m.aliens, alienName)
	
	}
	for _, inCityLinkName := range city.inLinks {
		inCityLink := m.cities[inCityLinkName]
		for linkDir, linkCityName := range inCityLink.outLinks {
			if linkCityName == city.name {
				delete(inCityLink.outLinks, linkDir)
				break
			}
		}
	}
	for _, outCityLinkName := range city.outLinks {
		outCityLink := m.cities[outCityLinkName]
		for linkDir, linkCityName:= range outCityLink.inLinks {
			if linkCityName == city.name {
				delete(outCityLink.inLinks, linkDir)
				break
			}
		}
	}
	delete(m.cities, city.name)
	return destroyedAliens
}
func (m *Map) ExecuteFights() {
	for _, alien := range m.aliens {
		occupiedCity:= alien.city
		city:= m.cities[occupiedCity]
		if len(city.alienOccupancy) == MaxOccupancy {
			destroyedAliens := m.destroyCity(city)
			log.Printf("%s has been destroyed by %s!", city.name, strings.Join(destroyedAliens, " and "))
		}
	}
}
func(m *Map) SeedAliens(n uint) {
	pq:= queue.NewPriorityQueue()
	for _, city := range m.cities {
		pq.Push(city)
	}
	seededAliens := uint(0)
	for seededAliens != n {
		city := pq.Pop().(*City)
	
	for i := 0; i<MaxOccupancy && seededAliens != n; i++ {
		alien := &Alien {
			name: 	fmt.Sprintf("alien%d", seededAliens+1),
			city: city.name,
		}
		city.alienOccupancy[alien.name] = alien
		m.aliens[alien.name] = alien
		seededAliens++ 
		}
	}
}
func (m *Map) String() (s string) {
	for _, city := range m.cities {
		aliens := make([]string, 0, len(city.alienOccupancy))
	for alien := range city.alienOccupancy {
		aliens = append(aliens, alien)
		}
		s+= fmt.Sprintf("{city:%s, outLinks:%s, inLinks:%s, alienOccupancy: [%s]}\n",
		city.name, city.outLinks, city.inLinks, strings.Join(aliens, " "))
	}
	return
}
