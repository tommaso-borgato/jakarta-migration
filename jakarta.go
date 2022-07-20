package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage %s <INPUT_FOLDER>\n", os.Args[0])
		os.Exit(3)
	}
	inputDirectory := os.Args[1]
	fmt.Printf("%s %s\n", os.Args[0], inputDirectory)

	err := filepath.Walk(inputDirectory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			processElem(path, info)
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

func processElem(path string, info os.FileInfo) {
	if !info.IsDir() {
		if strings.HasSuffix(path, ".java") {
			processJava(path, info)
		} else if strings.HasSuffix(path, "/pom.xml") {
			processPom(path, info)
		}
	}
}

func processJava(path string, info os.FileInfo) {
	fmt.Printf("\tprocessing java file %s ...\n", path)
	text := readFile(path)
	strings.ReplaceAll(text, "import javax.activation", "import jakarta.activation")
	strings.ReplaceAll(text, "import javax.annotation", "import jakarta.annotation")
	strings.ReplaceAll(text, "import javax.annotation.security", "import jakarta.annotation.security")
	strings.ReplaceAll(text, "import javax.annotation.sql", "import jakarta.annotation.sql")
	strings.ReplaceAll(text, "import javax.batch.api", "import jakarta.batch.api")
	/*strings.ReplaceAll(text, "import javax.batch.api.chunk", "import jakarta.batch.api.chunk")
	strings.ReplaceAll(text, "import javax.batch.api.chunk.listener", "import jakarta.batch.api.chunk.listener")
	strings.ReplaceAll(text, "import javax.batch.api.listener", "import jakarta.batch.api.listener")
	strings.ReplaceAll(text, "import javax.batch.api.partition", "import jakarta.batch.api.partition")*/
	strings.ReplaceAll(text, "import javax.batch.operations", "import jakarta.batch.operations")
	strings.ReplaceAll(text, "import javax.batch.runtime", "import jakarta.batch.runtime")
	strings.ReplaceAll(text, "import javax.batch.runtime.context", "import jakarta.batch.runtime.context")
	strings.ReplaceAll(text, "import javax.decorator", "import jakarta.decorator")
	strings.ReplaceAll(text, "import javax.ejb", "import jakarta.ejb")
	//strings.ReplaceAll(text, "import javax.ejb.embeddable", "import jakarta.ejb.embeddable")
	//strings.ReplaceAll(text, "import javax.ejb.spi", "import jakarta.ejb.spi")
	strings.ReplaceAll(text, "import javax.el", "import jakarta.el")
	strings.ReplaceAll(text, "import javax.enterprise.concurrent", "import jakarta.enterprise.concurrent")
	strings.ReplaceAll(text, "import javax.enterprise.context", "import jakarta.enterprise.context")
	strings.ReplaceAll(text, "import javax.enterprise.context.control", "import jakarta.enterprise.context.control")
	strings.ReplaceAll(text, "import javax.enterprise.context.spi", "import jakarta.enterprise.context.spi")
	strings.ReplaceAll(text, "import javax.enterprise.event", "import jakarta.enterprise.event")
	strings.ReplaceAll(text, "import javax.enterprise.inject", "import jakarta.enterprise.inject")
	strings.ReplaceAll(text, "import javax.enterprise.inject.literal", "import jakarta.enterprise.inject.literal")
	strings.ReplaceAll(text, "import javax.enterprise.inject.se", "import jakarta.enterprise.inject.se")
	strings.ReplaceAll(text, "import javax.enterprise.inject.spi", "import jakarta.enterprise.inject.spi")
	strings.ReplaceAll(text, "import javax.enterprise.inject.spi.configurator", "import jakarta.enterprise.inject.spi.configurator")
	strings.ReplaceAll(text, "import javax.enterprise.util", "import jakarta.enterprise.util")
	strings.ReplaceAll(text, "import javax.faces", "import jakarta.faces")
	/*strings.ReplaceAll(text, "import javax.faces.annotation", "import jakarta.faces.annotation")
	strings.ReplaceAll(text, "import javax.faces.application", "import jakarta.faces.application")
	strings.ReplaceAll(text, "import javax.faces.bean", "import jakarta.faces.bean")
	strings.ReplaceAll(text, "import javax.faces.component", "import jakarta.faces.component")
	strings.ReplaceAll(text, "import javax.faces.component.behavior", "import jakarta.faces.component.behavior")
	strings.ReplaceAll(text, "import javax.faces.component.html", "import jakarta.faces.component.html")
	strings.ReplaceAll(text, "import javax.faces.component.search", "import jakarta.faces.component.search")
	strings.ReplaceAll(text, "import javax.faces.component.visit", "import jakarta.faces.component.visit")
	strings.ReplaceAll(text, "import javax.faces.context", "import jakarta.faces.context")
	strings.ReplaceAll(text, "import javax.faces.convert", "import jakarta.faces.convert")
	strings.ReplaceAll(text, "import javax.faces.el", "import jakarta.faces.el")
	strings.ReplaceAll(text, "import javax.faces.event", "import jakarta.faces.event")
	strings.ReplaceAll(text, "import javax.faces.flow", "import jakarta.faces.flow")
	strings.ReplaceAll(text, "import javax.faces.flow.builder", "import jakarta.faces.flow.builder")
	strings.ReplaceAll(text, "import javax.faces.lifecycle", "import jakarta.faces.lifecycle")
	strings.ReplaceAll(text, "import javax.faces.model", "import jakarta.faces.model")
	strings.ReplaceAll(text, "import javax.faces.push", "import jakarta.faces.push")
	strings.ReplaceAll(text, "import javax.faces.render", "import jakarta.faces.render")
	strings.ReplaceAll(text, "import javax.faces.validator", "import jakarta.faces.validator")
	strings.ReplaceAll(text, "import javax.faces.view", "import jakarta.faces.view")
	strings.ReplaceAll(text, "import javax.faces.view.facelets", "import jakarta.faces.view.facelets")
	strings.ReplaceAll(text, "import javax.faces.webapp", "import jakarta.faces.webapp")*/
	strings.ReplaceAll(text, "import javax.inject", "import jakarta.inject")
	strings.ReplaceAll(text, "import javax.interceptor", "import jakarta.interceptor")
	strings.ReplaceAll(text, "import javax.jms", "import jakarta.jms")
	strings.ReplaceAll(text, "import javax.json", "import jakarta.json")
	/*strings.ReplaceAll(text, "import javax.json.bind", "import jakarta.json.bind")
	strings.ReplaceAll(text, "import javax.json.bind.adapter", "import jakarta.json.bind.adapter")
	strings.ReplaceAll(text, "import javax.json.bind.annotation", "import jakarta.json.bind.annotation")
	strings.ReplaceAll(text, "import javax.json.bind.config", "import jakarta.json.bind.config")
	strings.ReplaceAll(text, "import javax.json.bind.serializer", "import jakarta.json.bind.serializer")
	strings.ReplaceAll(text, "import javax.json.bind.spi", "import jakarta.json.bind.spi")
	strings.ReplaceAll(text, "import javax.json.spi", "import jakarta.json.spi")
	strings.ReplaceAll(text, "import javax.json.stream", "import jakarta.json.stream")*/
	strings.ReplaceAll(text, "import javax.jws", "import jakarta.jws")
	//strings.ReplaceAll(text, "import javax.jws.soap", "import jakarta.jws.soap")
	strings.ReplaceAll(text, "import javax.mail", "import jakarta.mail")
	/*strings.ReplaceAll(text, "import javax.mail.event", "import jakarta.mail.event")
	strings.ReplaceAll(text, "import javax.mail.internet", "import jakarta.mail.internet")
	strings.ReplaceAll(text, "import javax.mail.search", "import jakarta.mail.search")
	strings.ReplaceAll(text, "import javax.mail.util", "import jakarta.mail.util")*/
	strings.ReplaceAll(text, "import javax.persistence", "import jakarta.persistence")
	/*strings.ReplaceAll(text, "import javax.persistence.criteria", "import jakarta.persistence.criteria")
	strings.ReplaceAll(text, "import javax.persistence.metamodel", "import jakarta.persistence.metamodel")
	strings.ReplaceAll(text, "import javax.persistence.spi", "import jakarta.persistence.spi")*/
	strings.ReplaceAll(text, "import javax.resource", "import jakarta.resource")
	//strings.ReplaceAll(text, "import javax.resource.cci", "import jakarta.resource.cci")
	strings.ReplaceAll(text, "import javax.resource.spi", "import jakarta.resource.spi")
	/*strings.ReplaceAll(text, "import javax.resource.spi.endpoint", "import jakarta.resource.spi.endpoint")
	strings.ReplaceAll(text, "import javax.resource.spi.security", "import jakarta.resource.spi.security")
	strings.ReplaceAll(text, "import javax.resource.spi.work", "import jakarta.resource.spi.work")*/
	strings.ReplaceAll(text, "import javax.security.auth.message", "import jakarta.security.auth.message")
	/*strings.ReplaceAll(text, "import javax.security.auth.message.callback", "import jakarta.security.auth.message.callback")
	strings.ReplaceAll(text, "import javax.security.auth.message.config", "import jakarta.security.auth.message.config")
	strings.ReplaceAll(text, "import javax.security.auth.message.module", "import jakarta.security.auth.message.module")*/
	strings.ReplaceAll(text, "import javax.security.enterprise", "import jakarta.security.enterprise")
	strings.ReplaceAll(text, "import javax.security.enterprise.authentication.mechanism.http", "import jakarta.security.enterprise.authentication.mechanism.http")
	strings.ReplaceAll(text, "import javax.security.enterprise.credential", "import jakarta.security.enterprise.credential")
	strings.ReplaceAll(text, "import javax.security.enterprise.identitystore", "import jakarta.security.enterprise.identitystore")
	strings.ReplaceAll(text, "import javax.security.jacc", "import jakarta.security.jacc")
	strings.ReplaceAll(text, "import javax.servlet", "import jakarta.servlet")
	/*strings.ReplaceAll(text, "import javax.servlet.annotation", "import jakarta.servlet.annotation")
	strings.ReplaceAll(text, "import javax.servlet.descriptor", "import jakarta.servlet.descriptor")
	strings.ReplaceAll(text, "import javax.servlet.http", "import jakarta.servlet.http")
	strings.ReplaceAll(text, "import javax.servlet.jsp", "import jakarta.servlet.jsp")
	strings.ReplaceAll(text, "import javax.servlet.jsp.el", "import jakarta.servlet.jsp.el")
	strings.ReplaceAll(text, "import javax.servlet.jsp.jstl.core", "import jakarta.servlet.jsp.jstl.core")
	strings.ReplaceAll(text, "import javax.servlet.jsp.jstl.fmt", "import jakarta.servlet.jsp.jstl.fmt")
	strings.ReplaceAll(text, "import javax.servlet.jsp.jstl.sql", "import jakarta.servlet.jsp.jstl.sql")
	strings.ReplaceAll(text, "import javax.servlet.jsp.jstl.tlv", "import jakarta.servlet.jsp.jstl.tlv")
	strings.ReplaceAll(text, "import javax.servlet.jsp.tagext", "import jakarta.servlet.jsp.tagext")*/
	strings.ReplaceAll(text, "import javax.transaction", "import jakarta.transaction")
	strings.ReplaceAll(text, "import javax.validation", "import jakarta.validation")
	/*strings.ReplaceAll(text, "import javax.validation.bootstrap", "import jakarta.validation.bootstrap")
	strings.ReplaceAll(text, "import javax.validation.constraints", "import jakarta.validation.constraints")
	strings.ReplaceAll(text, "import javax.validation.constraintvalidation", "import jakarta.validation.constraintvalidation")
	strings.ReplaceAll(text, "import javax.validation.executable", "import jakarta.validation.executable")
	strings.ReplaceAll(text, "import javax.validation.groups", "import jakarta.validation.groups")
	strings.ReplaceAll(text, "import javax.validation.metadata", "import jakarta.validation.metadata")
	strings.ReplaceAll(text, "import javax.validation.spi", "import jakarta.validation.spi")
	strings.ReplaceAll(text, "import javax.validation.valueextraction", "import jakarta.validation.valueextraction")*/
	strings.ReplaceAll(text, "import javax.websocket", "import jakarta.websocket")
	//strings.ReplaceAll(text, "import javax.websocket.server", "import jakarta.websocket.server")
	strings.ReplaceAll(text, "import javax.ws.rs", "import jakarta.ws.rs")
	/*strings.ReplaceAll(text, "import javax.ws.rs.client", "import jakarta.ws.rs.client")
	strings.ReplaceAll(text, "import javax.ws.rs.container", "import jakarta.ws.rs.container")
	strings.ReplaceAll(text, "import javax.ws.rs.core", "import jakarta.ws.rs.core")
	strings.ReplaceAll(text, "import javax.ws.rs.ext", "import jakarta.ws.rs.ext")
	strings.ReplaceAll(text, "import javax.ws.rs.sse", "import jakarta.ws.rs.sse")*/
	strings.ReplaceAll(text, "import javax.xml.bind", "import jakarta.xml.bind")
	/*strings.ReplaceAll(text, "import javax.xml.bind.annotation", "import jakarta.xml.bind.annotation")
	strings.ReplaceAll(text, "import javax.xml.bind.annotation.adapters", "import jakarta.xml.bind.annotation.adapters")
	strings.ReplaceAll(text, "import javax.xml.bind.attachment", "import jakarta.xml.bind.attachment")
	strings.ReplaceAll(text, "import javax.xml.bind.helpers", "import jakarta.xml.bind.helpers")
	strings.ReplaceAll(text, "import javax.xml.bind.util", "import jakarta.xml.bind.util")*/
	strings.ReplaceAll(text, "import javax.xml.soap", "import jakarta.xml.soap")
	strings.ReplaceAll(text, "import javax.xml.ws", "import jakarta.xml.ws")
	/*strings.ReplaceAll(text, "import javax.xml.ws.handler", "import jakarta.xml.ws.handler")
	strings.ReplaceAll(text, "import javax.xml.ws.handler.soap", "import jakarta.xml.ws.handler.soap")
	strings.ReplaceAll(text, "import javax.xml.ws.http", "import jakarta.xml.ws.http")
	strings.ReplaceAll(text, "import javax.xml.ws.soap", "import jakarta.xml.ws.soap")
	strings.ReplaceAll(text, "import javax.xml.ws.spi", "import jakarta.xml.ws.spi")
	strings.ReplaceAll(text, "import javax.xml.ws.spi.http", "import jakarta.xml.ws.spi.http")
	strings.ReplaceAll(text, "import javax.xml.ws.wsaddressing", "import jakarta.xml.ws.wsaddressing")*/
	fmt.Printf("%s", text)
}

func processPom(path string, info os.FileInfo) {
	fmt.Printf("\tprocessing pom.xml file %s ...\n", path)
	//text := readFile(path)
}

func readFile(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	return text
}
