package tiled

import (
	"encoding/json"
	"io/ioutil"
)

/*
	Based on the JSON Map Format for Tiled 1.2.
	https://doc.mapeditor.org/en/stable/reference/json-map-format/

	TODO:
		- This needs to be extensively tested. You can help by e.g.
		  sending me your Tiled files (as JSON) in order to give me
		  testing material.
		- Comment all fields with the description in the link above.
		  Textually separate optional fields and describe them as such.
*/

// Parse takes a Tiled map file exported as JSON and parses it into
// a Golang data structure.
func Parse(filename string) (*Map, error) {
	m := &Map{}
	jdata, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jdata, m)
	return m, err
}

// Map describes a Tiled map.
type Map struct {
	BackgroundColor string     `json:"backgroundcolor"`
	Height          int        `json:"height"`
	HexSideLength   int        `json:"hexsidelength"`
	Infinite        bool       `json:"infinite"`
	Layers          []Layer    `json:"layers"`
	NextLayerID     int        `json:"nextlayerid"`
	NextObjectID    int        `json:"nextobjectid"`
	Orientation     string     `json:"orientation"`
	Properties      []Property `json:"properties"`
	RenderOrder     string     `json:"renderorder"`
	StaggerAxis     string     `json:"staggeraxis"`
	StaggerIndex    string     `json:"staggerindex"`
	TiledVersion    string     `json:"tiledversion"`
	TileHeight      int        `json:"tileheight"`
	Tilesets        []Tileset  `json:"tilesets"`
	TileWidth       int        `json:"tilewidth"`
	Type            string     `json:"type"`
	Version         float64    `json:"version"`
	Width           int        `json:"width"`
}

type Property struct {
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type Layer struct {
	// Common
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Visible    bool       `json:"visible"`
	Width      int        `json:"width"`
	Height     int        `json:"height"`
	X          int        `json:"x"`
	Y          int        `json:"y"`
	Offsetx    float64    `json:"offsetx"`
	Offsety    float64    `json:"offsety"`
	Opacity    float64    `json:"opacity"`
	Properties []Property `json:"properties"`

	// TileLayer only
	Chunks      []Chunk `json:"chunks"`
	Compression string  `json:"compression"`
	// Array or string. Array of unsigned int (GIDs) or base64-encoded data.
	Data     interface{} `json:"data"`
	Encoding string      `json:"encoding"`

	// ObjectGroup only
	Draworder string `json:"draworder"`

	// Group only
	Layers  []Layer  `json:"layers"`
	Objects []Object `json:"objects"`

	// ImageLayer only
	Image            string `json:"image"`
	Transparentcolor string `json:"transparentcolor"`
}

// Chunk is used to store the tile layer data for infinite maps.
type Chunk struct {
	Data   interface{} `json:"data"` // Array of unsigned int (GIDs) or base64-encoded data
	Height int         `json:"height"`
	Width  int         `json:"width"`
	X      int         `json:"x"`
	Y      int         `json:"y"`
}

type Object struct {
	ID         int                    `json:"id"`
	GID        int                    `json:"gid"`
	Name       string                 `json:"name"`
	Type       string                 `json:"type"`
	X          float64                `json:"x"`
	Y          float64                `json:"y"`
	Width      float64                `json:"width"`
	Height     float64                `json:"height"`
	Visible    bool                   `json:"visible"`
	Ellipse    bool                   `json:"ellipse"`
	Point      bool                   `json:"point"`
	Polygon    []Coordinate           `json:"polygon"`
	Polyline   []Coordinate           `json:"polyline"`
	Properties []Property             `json:"properties"`
	Rotation   float64                `json:"rotation"`
	Template   string                 `json:"template"`
	Text       map[string]interface{} `json:"text"`
}

type Coordinate struct {
	X, Y float64
}

type Tileset struct {
	FirstGID         int        `json:"firstgid"`
	Name             string     `json:"name"`
	Type             string     `json:"type"`
	Columns          int        `json:"columns"`
	Grid             Grid       `json:"grid"`
	Image            string     `json:"image"`
	ImageHeight      int        `json:"imageheight"`
	ImageWidth       int        `json:"imagewidth"`
	Properties       []Property `json:"properties"`
	Margin           int        `json:"margin"`
	Spacing          int        `json:"spacing"`
	TileCount        int        `json:"tilecount"`
	TileWidth        int        `json:"tilewidth"`
	TileHeight       int        `json:"tileheight"`
	TransparentColor string     `json:"transparentcolor"`
	TileOffset       struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"tileoffset"`
	Terrains []Terrain `json:"terrains"`
	Tiles    []Tile    `json:"tiles"`
	WangSets []WangSet `json:"wangsets"`
}

type Grid struct {
	Orientation string `json:"orientation"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
}

type Tile struct {
	ID          int        `json:"id"`
	Type        string     `json:"type"`
	Properties  []Property `json:"properties"`
	Animation   []Frame    `json:"animation"`
	Terrain     []int      `json:"terrain"`
	Image       string     `json:"image"`
	ImageHeight int        `json:"imageheight"`
	ImageWidth  int        `json:"imagewidth"`
	ObjectGroup Layer      `json:"objectgroup"`
}

type Frame struct {
	Duration int `json:"duration"`
	TileID   int `json:"tileid"`
}

type Terrain struct {
	Name string `json:"name"`
	Tile int    `json:"tile"`
}

type WangSet struct {
	CornerColors []WangColor `json:"cornercolors"`
	EdgeColors   []WangColor `json:"edgecolors"`
	Name         string      `json:"name"`
	Tile         int         `json:"tile"`
	WangTiles    []WangTile  `json:"wangtiles"`
}

type WangColor struct {
	Color       string  `json:"color"`
	Name        string  `json:"name"`
	Probability float64 `json:"probability"`
	Tile        int     `json:"tile"`
}

type WangTile struct {
	TileID int     `json:"tileid"`
	WangID [8]byte `json:"wangid"`
	Dflip  bool    `json:"dflip"`
	Hflip  bool    `json:"hflip"`
	Vflip  bool    `json:"vflip"`
}

type ObjectTemplate struct {
	Type    string  `json:"type"`
	Tileset Tileset `json:"tileset"`
	Object  Object  `json:"object"`
}
