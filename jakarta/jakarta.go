package jakarta

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func ProcessDirectory(inputDirectory string) {
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
	//fmt.Printf("\tprocessing java file %s ...\n", path)
	text := readFile(path)
	finaltext := strings.ReplaceAll(text, "import javax.activation", "import jakarta.activation")
	finaltext = strings.ReplaceAll(finaltext, "import javax.annotation", "import jakarta.annotation")
	finaltext = strings.ReplaceAll(finaltext, "import javax.annotation.security", "import jakarta.annotation.security")
	finaltext = strings.ReplaceAll(finaltext, "import javax.annotation.sql", "import jakarta.annotation.sql")
	finaltext = strings.ReplaceAll(finaltext, "import javax.batch.api", "import jakarta.batch.api")
	/*strings.ReplaceAll(finaltext, "import javax.batch.api.chunk", "import jakarta.batch.api.chunk")
	strings.ReplaceAll(finaltext, "import javax.batch.api.chunk.listener", "import jakarta.batch.api.chunk.listener")
	strings.ReplaceAll(finaltext, "import javax.batch.api.listener", "import jakarta.batch.api.listener")
	strings.ReplaceAll(finaltext, "import javax.batch.api.partition", "import jakarta.batch.api.partition")*/
	finaltext = strings.ReplaceAll(finaltext, "import javax.batch.operations", "import jakarta.batch.operations")
	finaltext = strings.ReplaceAll(finaltext, "import javax.batch.runtime", "import jakarta.batch.runtime")
	finaltext = strings.ReplaceAll(finaltext, "import javax.batch.runtime.context", "import jakarta.batch.runtime.context")
	finaltext = strings.ReplaceAll(finaltext, "import javax.decorator", "import jakarta.decorator")
	finaltext = strings.ReplaceAll(finaltext, "import javax.ejb", "import jakarta.ejb")
	//strings.ReplaceAll(finaltext, "import javax.ejb.embeddable", "import jakarta.ejb.embeddable")
	//strings.ReplaceAll(finaltext, "import javax.ejb.spi", "import jakarta.ejb.spi")
	finaltext = strings.ReplaceAll(finaltext, "import javax.el", "import jakarta.el")
	finaltext = strings.ReplaceAll(finaltext, "import javax.enterprise.concurrent", "import jakarta.enterprise.concurrent")
	finaltext = strings.ReplaceAll(finaltext, "import javax.enterprise.context", "import jakarta.enterprise.context")
	finaltext = strings.ReplaceAll(finaltext, "import javax.enterprise.context.control", "import jakarta.enterprise.context.control")
	finaltext = strings.ReplaceAll(finaltext, "import javax.enterprise.context.spi", "import jakarta.enterprise.context.spi")
	finaltext = strings.ReplaceAll(finaltext, "import javax.enterprise.event", "import jakarta.enterprise.event")
	finaltext = strings.ReplaceAll(finaltext, "import javax.enterprise.inject", "import jakarta.enterprise.inject")
	finaltext = strings.ReplaceAll(finaltext, "import javax.enterprise.inject.literal", "import jakarta.enterprise.inject.literal")
	finaltext = strings.ReplaceAll(finaltext, "import javax.enterprise.inject.se", "import jakarta.enterprise.inject.se")
	finaltext = strings.ReplaceAll(finaltext, "import javax.enterprise.inject.spi", "import jakarta.enterprise.inject.spi")
	finaltext = strings.ReplaceAll(finaltext, "import javax.enterprise.inject.spi.configurator", "import jakarta.enterprise.inject.spi.configurator")
	finaltext = strings.ReplaceAll(finaltext, "import javax.enterprise.util", "import jakarta.enterprise.util")
	finaltext = strings.ReplaceAll(finaltext, "import javax.faces", "import jakarta.faces")
	/*strings.ReplaceAll(finaltext, "import javax.faces.annotation", "import jakarta.faces.annotation")
	strings.ReplaceAll(finaltext, "import javax.faces.application", "import jakarta.faces.application")
	strings.ReplaceAll(finaltext, "import javax.faces.bean", "import jakarta.faces.bean")
	strings.ReplaceAll(finaltext, "import javax.faces.component", "import jakarta.faces.component")
	strings.ReplaceAll(finaltext, "import javax.faces.component.behavior", "import jakarta.faces.component.behavior")
	strings.ReplaceAll(finaltext, "import javax.faces.component.html", "import jakarta.faces.component.html")
	strings.ReplaceAll(finaltext, "import javax.faces.component.search", "import jakarta.faces.component.search")
	strings.ReplaceAll(finaltext, "import javax.faces.component.visit", "import jakarta.faces.component.visit")
	strings.ReplaceAll(finaltext, "import javax.faces.context", "import jakarta.faces.context")
	strings.ReplaceAll(finaltext, "import javax.faces.convert", "import jakarta.faces.convert")
	strings.ReplaceAll(finaltext, "import javax.faces.el", "import jakarta.faces.el")
	strings.ReplaceAll(finaltext, "import javax.faces.event", "import jakarta.faces.event")
	strings.ReplaceAll(finaltext, "import javax.faces.flow", "import jakarta.faces.flow")
	strings.ReplaceAll(finaltext, "import javax.faces.flow.builder", "import jakarta.faces.flow.builder")
	strings.ReplaceAll(finaltext, "import javax.faces.lifecycle", "import jakarta.faces.lifecycle")
	strings.ReplaceAll(finaltext, "import javax.faces.model", "import jakarta.faces.model")
	strings.ReplaceAll(finaltext, "import javax.faces.push", "import jakarta.faces.push")
	strings.ReplaceAll(finaltext, "import javax.faces.render", "import jakarta.faces.render")
	strings.ReplaceAll(finaltext, "import javax.faces.validator", "import jakarta.faces.validator")
	strings.ReplaceAll(finaltext, "import javax.faces.view", "import jakarta.faces.view")
	strings.ReplaceAll(finaltext, "import javax.faces.view.facelets", "import jakarta.faces.view.facelets")
	strings.ReplaceAll(finaltext, "import javax.faces.webapp", "import jakarta.faces.webapp")*/
	finaltext = strings.ReplaceAll(finaltext, "import javax.inject", "import jakarta.inject")
	finaltext = strings.ReplaceAll(finaltext, "import javax.interceptor", "import jakarta.interceptor")
	finaltext = strings.ReplaceAll(finaltext, "import javax.jms", "import jakarta.jms")
	finaltext = strings.ReplaceAll(finaltext, "import javax.json", "import jakarta.json")
	/*strings.ReplaceAll(finaltext, "import javax.json.bind", "import jakarta.json.bind")
	strings.ReplaceAll(finaltext, "import javax.json.bind.adapter", "import jakarta.json.bind.adapter")
	strings.ReplaceAll(finaltext, "import javax.json.bind.annotation", "import jakarta.json.bind.annotation")
	strings.ReplaceAll(finaltext, "import javax.json.bind.config", "import jakarta.json.bind.config")
	strings.ReplaceAll(finaltext, "import javax.json.bind.serializer", "import jakarta.json.bind.serializer")
	strings.ReplaceAll(finaltext, "import javax.json.bind.spi", "import jakarta.json.bind.spi")
	strings.ReplaceAll(finaltext, "import javax.json.spi", "import jakarta.json.spi")
	strings.ReplaceAll(finaltext, "import javax.json.stream", "import jakarta.json.stream")*/
	finaltext = strings.ReplaceAll(finaltext, "import javax.jws", "import jakarta.jws")
	//strings.ReplaceAll(finaltext, "import javax.jws.soap", "import jakarta.jws.soap")
	finaltext = strings.ReplaceAll(finaltext, "import javax.mail", "import jakarta.mail")
	/*strings.ReplaceAll(finaltext, "import javax.mail.event", "import jakarta.mail.event")
	strings.ReplaceAll(finaltext, "import javax.mail.internet", "import jakarta.mail.internet")
	strings.ReplaceAll(finaltext, "import javax.mail.search", "import jakarta.mail.search")
	strings.ReplaceAll(finaltext, "import javax.mail.util", "import jakarta.mail.util")*/
	finaltext = strings.ReplaceAll(finaltext, "import javax.persistence", "import jakarta.persistence")
	/*strings.ReplaceAll(finaltext, "import javax.persistence.criteria", "import jakarta.persistence.criteria")
	strings.ReplaceAll(finaltext, "import javax.persistence.metamodel", "import jakarta.persistence.metamodel")
	strings.ReplaceAll(finaltext, "import javax.persistence.spi", "import jakarta.persistence.spi")*/
	finaltext = strings.ReplaceAll(finaltext, "import javax.resource", "import jakarta.resource")
	//strings.ReplaceAll(finaltext, "import javax.resource.cci", "import jakarta.resource.cci")
	finaltext = strings.ReplaceAll(finaltext, "import javax.resource.spi", "import jakarta.resource.spi")
	/*strings.ReplaceAll(finaltext, "import javax.resource.spi.endpoint", "import jakarta.resource.spi.endpoint")
	strings.ReplaceAll(finaltext, "import javax.resource.spi.security", "import jakarta.resource.spi.security")
	strings.ReplaceAll(finaltext, "import javax.resource.spi.work", "import jakarta.resource.spi.work")*/
	finaltext = strings.ReplaceAll(finaltext, "import javax.security.auth.message", "import jakarta.security.auth.message")
	/*strings.ReplaceAll(finaltext, "import javax.security.auth.message.callback", "import jakarta.security.auth.message.callback")
	strings.ReplaceAll(finaltext, "import javax.security.auth.message.config", "import jakarta.security.auth.message.config")
	strings.ReplaceAll(finaltext, "import javax.security.auth.message.module", "import jakarta.security.auth.message.module")*/
	finaltext = strings.ReplaceAll(finaltext, "import javax.security.enterprise", "import jakarta.security.enterprise")
	finaltext = strings.ReplaceAll(finaltext, "import javax.security.enterprise.authentication.mechanism.http", "import jakarta.security.enterprise.authentication.mechanism.http")
	finaltext = strings.ReplaceAll(finaltext, "import javax.security.enterprise.credential", "import jakarta.security.enterprise.credential")
	finaltext = strings.ReplaceAll(finaltext, "import javax.security.enterprise.identitystore", "import jakarta.security.enterprise.identitystore")
	finaltext = strings.ReplaceAll(finaltext, "import javax.security.jacc", "import jakarta.security.jacc")
	finaltext = strings.ReplaceAll(finaltext, "import javax.servlet", "import jakarta.servlet")
	/*strings.ReplaceAll(finaltext, "import javax.servlet.annotation", "import jakarta.servlet.annotation")
	strings.ReplaceAll(finaltext, "import javax.servlet.descriptor", "import jakarta.servlet.descriptor")
	strings.ReplaceAll(finaltext, "import javax.servlet.http", "import jakarta.servlet.http")
	strings.ReplaceAll(finaltext, "import javax.servlet.jsp", "import jakarta.servlet.jsp")
	strings.ReplaceAll(finaltext, "import javax.servlet.jsp.el", "import jakarta.servlet.jsp.el")
	strings.ReplaceAll(finaltext, "import javax.servlet.jsp.jstl.core", "import jakarta.servlet.jsp.jstl.core")
	strings.ReplaceAll(finaltext, "import javax.servlet.jsp.jstl.fmt", "import jakarta.servlet.jsp.jstl.fmt")
	strings.ReplaceAll(finaltext, "import javax.servlet.jsp.jstl.sql", "import jakarta.servlet.jsp.jstl.sql")
	strings.ReplaceAll(finaltext, "import javax.servlet.jsp.jstl.tlv", "import jakarta.servlet.jsp.jstl.tlv")
	strings.ReplaceAll(finaltext, "import javax.servlet.jsp.tagext", "import jakarta.servlet.jsp.tagext")*/
	finaltext = strings.ReplaceAll(finaltext, "import javax.transaction", "import jakarta.transaction")
	finaltext = strings.ReplaceAll(finaltext, "import javax.validation", "import jakarta.validation")
	/*strings.ReplaceAll(finaltext, "import javax.validation.bootstrap", "import jakarta.validation.bootstrap")
	strings.ReplaceAll(finaltext, "import javax.validation.constraints", "import jakarta.validation.constraints")
	strings.ReplaceAll(finaltext, "import javax.validation.constraintvalidation", "import jakarta.validation.constraintvalidation")
	strings.ReplaceAll(finaltext, "import javax.validation.executable", "import jakarta.validation.executable")
	strings.ReplaceAll(finaltext, "import javax.validation.groups", "import jakarta.validation.groups")
	strings.ReplaceAll(finaltext, "import javax.validation.metadata", "import jakarta.validation.metadata")
	strings.ReplaceAll(finaltext, "import javax.validation.spi", "import jakarta.validation.spi")
	strings.ReplaceAll(finaltext, "import javax.validation.valueextraction", "import jakarta.validation.valueextraction")*/
	finaltext = strings.ReplaceAll(finaltext, "import javax.websocket", "import jakarta.websocket")
	//strings.ReplaceAll(finaltext, "import javax.websocket.server", "import jakarta.websocket.server")
	finaltext = strings.ReplaceAll(finaltext, "import javax.ws.rs", "import jakarta.ws.rs")
	/*strings.ReplaceAll(finaltext, "import javax.ws.rs.client", "import jakarta.ws.rs.client")
	strings.ReplaceAll(finaltext, "import javax.ws.rs.container", "import jakarta.ws.rs.container")
	strings.ReplaceAll(finaltext, "import javax.ws.rs.core", "import jakarta.ws.rs.core")
	strings.ReplaceAll(finaltext, "import javax.ws.rs.ext", "import jakarta.ws.rs.ext")
	strings.ReplaceAll(finaltext, "import javax.ws.rs.sse", "import jakarta.ws.rs.sse")*/
	finaltext = strings.ReplaceAll(finaltext, "import javax.xml.bind", "import jakarta.xml.bind")
	/*strings.ReplaceAll(finaltext, "import javax.xml.bind.annotation", "import jakarta.xml.bind.annotation")
	strings.ReplaceAll(finaltext, "import javax.xml.bind.annotation.adapters", "import jakarta.xml.bind.annotation.adapters")
	strings.ReplaceAll(finaltext, "import javax.xml.bind.attachment", "import jakarta.xml.bind.attachment")
	strings.ReplaceAll(finaltext, "import javax.xml.bind.helpers", "import jakarta.xml.bind.helpers")
	strings.ReplaceAll(finaltext, "import javax.xml.bind.util", "import jakarta.xml.bind.util")*/
	finaltext = strings.ReplaceAll(finaltext, "import javax.xml.soap", "import jakarta.xml.soap")
	finaltext = strings.ReplaceAll(finaltext, "import javax.xml.ws", "import jakarta.xml.ws")
	/*strings.ReplaceAll(finaltext, "import javax.xml.ws.handler", "import jakarta.xml.ws.handler")
	strings.ReplaceAll(finaltext, "import javax.xml.ws.handler.soap", "import jakarta.xml.ws.handler.soap")
	strings.ReplaceAll(finaltext, "import javax.xml.ws.http", "import jakarta.xml.ws.http")
	strings.ReplaceAll(finaltext, "import javax.xml.ws.soap", "import jakarta.xml.ws.soap")
	strings.ReplaceAll(finaltext, "import javax.xml.ws.spi", "import jakarta.xml.ws.spi")
	strings.ReplaceAll(finaltext, "import javax.xml.ws.spi.http", "import jakarta.xml.ws.spi.http")
	strings.ReplaceAll(finaltext, "import javax.xml.ws.wsaddressing", "import jakarta.xml.ws.wsaddressing")*/
	//fmt.Printf("%s", finaltext)
	if strings.Compare(text, finaltext) != 0 {
		fmt.Printf("Java file %s was modified.\n", path)
		writeFile(path, finaltext)
	}
}

func processPom(path string, info os.FileInfo) {
	//fmt.Printf("\tprocessing pom.xml file %s ...\n", path)
	//text := readFile(path)

	fmt.Printf("\nmvn dependency:tree -f %s:\n\n", path)
	cmd := exec.Command("mvn", "dependency:tree", "-f", path)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(fmt.Sprintf("Error setting up mvn with pom file %s.: ", path), err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(fmt.Sprintf("Error running mvn with pom file %s.: ", path), err)
	}

	data, err := ioutil.ReadAll(stdout)

	if err != nil {
		fmt.Printf("%s\n", string(data))
		log.Fatal(fmt.Sprintf("Error reading output from mvn with pom file %s.: ", path), err)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Printf("%s\n", string(data))
		log.Fatal(fmt.Sprintf("Error waiting output from mvn with pom file %s.: ", path), err)
	}

	//fmt.Printf("%s\n", string(data))

	GAVS := [...]string{
		"javax.ejb",
		"javax.activation",
		"javax.annotation",
		"javax.annotation.security",
		"javax.annotation.sql",
		"javax.batch.api",
		"javax.batch.operations",
		"javax.batch.runtime",
		"javax.decorator",
		"javax.el",
		"javax.enterprise", // added by me, not in javadoc
		"javax.enterprise.concurrent",
		"javax.enterprise.context",
		"javax.enterprise.context.control",
		"javax.enterprise.context.spi",
		"javax.enterprise.event",
		"javax.enterprise.inject",
		"javax.enterprise.inject.literal",
		"javax.enterprise.inject.se",
		"javax.enterprise.inject.spi",
		"javax.enterprise.inject.spi.configurator",
		"javax.enterprise.util",
		"javax.faces",
		"javax.inject",
		"javax.interceptor",
		"javax.jms",
		"javax.json",
		"javax.jws",
		"javax.mail",
		"javax.persistence",
		"javax.resource",
		"javax.security.auth.message",
		"javax.security.enterprise",
		"javax.security.jacc",
		"javax.servlet",
		"javax.transaction",
		"javax.validation",
		"javax.websocket",
		"javax.ws.rs",
		"javax.xml.bind",
		"javax.xml.soap",
		"javax.xml.ws",
	}

	for _, value := range GAVS {
		r := regexp.MustCompile(fmt.Sprintf("[^ ]*%s.*:", value))
		matches := r.FindAllString(string(data), -1)
		for _, v := range matches {
			//fmt.Printf("Dependency to remove: %s\n", strings.TrimSuffix(v, ":"))

			cmd := exec.Command("mvn", "dependency:tree", "-f", path, fmt.Sprintf("-Dincludes=%s", strings.TrimSuffix(v, ":")))

			stdout, err := cmd.StdoutPipe()

			if err != nil {
				log.Fatal(fmt.Sprintf("Error setting up mvn with pom file %s.: ", path), err)
			}

			if err := cmd.Start(); err != nil {
				log.Fatal(fmt.Sprintf("Error running mvn with pom file %s.: ", path), err)
			}

			datainc, err := ioutil.ReadAll(stdout)

			if err != nil {
				fmt.Printf("%s\n", string(datainc))
				log.Fatal(fmt.Sprintf("Error reading output from mvn with pom file %s.: ", path), err)
			}

			if err := cmd.Wait(); err != nil {
				fmt.Printf("%s\n", string(datainc))
				log.Fatal(fmt.Sprintf("Error waiting output from mvn with pom file %s.: ", path), err)
			}

			fmt.Printf("#################### Depemdency to remove/replace: ####################\n%s\n\n", string(datainc))
		}
	}
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

func writeFile(path string, text string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.WriteString(text)
	if err != nil {
		log.Fatal(err)
	}
}
