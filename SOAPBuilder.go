package gosoap

import (
	"github.com/beevik/etree"
	"log"
)

type SoapMessage string

func NewSOAP(headContent []*etree.Element, bodyContent []*etree.Element, namespaces map[string]string) SoapMessage {
	doc := buildSoapRoot()
	doc.IndentTabs()

	res, _ := doc.WriteToString();

	return  SoapMessage(res)
}

func (msg SoapMessage) String() string {
	return string(msg)
}

func (msg *SoapMessage) AddBodyContent(element *etree.Element)  {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(msg.String()); err != nil {
		log.Println(err.Error())
	}
	//doc.FindElement("./Envelope/Body").AddChild(element)
	bodyTag := doc.Root().SelectElement("Body")
	bodyTag.AddChild(element)

	doc.IndentTabs()
	res, _ := doc.WriteToString()

	*msg = SoapMessage(res)
}

func (msg *SoapMessage) AddBodyContents(elements []*etree.Element) {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(msg.String()); err != nil {
		log.Println(err.Error())
	}

	bodyTag := doc.Root().SelectElement("Body")

	if len(elements) != 0 {
		for _, j := range elements {
			bodyTag.AddChild(j)
		}
	}

	doc.IndentTabs()
	res, _ := doc.WriteToString()

	*msg = SoapMessage(res)
}

func (msg *SoapMessage) AddHeaderContent(element *etree.Element) {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(msg.String()); err != nil {
		log.Println(err.Error())
	}
	//doc.FindElement("./Envelope/Body").AddChild(element)
	bodyTag := doc.Root().SelectElement("Header")
	bodyTag.AddChild(element)

	doc.IndentTabs()
	res, _ := doc.WriteToString()

	*msg = SoapMessage(res)
}

func (msg *SoapMessage) AddHeaderContents(elements []*etree.Element) {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(msg.String()); err != nil {
		log.Println(err.Error())
	}

	headerTag := doc.Root().SelectElement("Header")

	if len(elements) != 0 {
		for _, j := range elements {
			headerTag.AddChild(j)
		}
	}

	doc.IndentTabs()
	res, _ := doc.WriteToString()

	*msg = SoapMessage(res)
}

func (msg *SoapMessage) AddRootNamespace(key, value string) {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(msg.String()); err != nil {
		log.Println(err.Error())
	}

	doc.Root().CreateAttr("xmlns:" + key, value)

	doc.IndentTabs()
	res, _ := doc.WriteToString()

	*msg = SoapMessage(res)
}

func (msg *SoapMessage) AddRootNamespaces(namespaces map[string]string) {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(msg.String()); err != nil {
		log.Println(err.Error())
	}

	for key, value := range namespaces {
		doc.CreateAttr("xmlns:" + key, value)
	}

	doc.IndentTabs()
	res, _ := doc.WriteToString()

	*msg = SoapMessage(res)
}


func buildSoapRoot() *etree.Document {
	doc := etree.NewDocument()

	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	env := doc.CreateElement("Envelope")
	env.CreateElement("Header")
	env.CreateElement("Body")

	env.CreateAttr("xmlns", "http://www.w3.org/2003/05/soap-envelope")

	return doc
}