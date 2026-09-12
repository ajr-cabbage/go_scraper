package main

import (
	"testing"
)

func TestGetHeadingFromHTMLBasic(t *testing.T) {
	inputBody := "<html><body><h1>Test Title</h1></body></html>"
	actual := getHeadingFromHTML(inputBody)
	expected := "Test Title"

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetHeadingFromHTMLH2(t *testing.T) {
	inputBody := "<html><body><h2>Test Title</h2></body></html>"
	actual := getHeadingFromHTML(inputBody)
	expected := "Test Title"

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetHeadingFromHTMLNoHeading(t *testing.T) {
	inputBody := "<html><body><p>No Title Here</p></body></html>"
	actual := getHeadingFromHTML(inputBody)
	expected := ""

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {
	inputBody := `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := "Main paragraph."

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetFirstParagraphFromHTMLOutsideFallback(t *testing.T) {
	inputBody := `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<a>Main paragraph.</a>
		</main>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := "Outside paragraph."

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetFirstParagraphFromHTMLNoP(t *testing.T) {
	inputBody := `<html><body>
		<b>Outside paragraph.</b>
		<main>
			<b>Main paragraph.</b>
		</main>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := ""

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetFirstParagraphFromHTMLManyP(t *testing.T) {
	inputBody := `<html><body>
		<p>Outside paragraph 1.</p>
		<p>Outside paragraph 2.</p>
		<p>Outside paragraph 3.</p>
		<p>Outside paragraph 4.</p>
		<main>
			<p>Main paragraph 1.</p>
			<p>Main paragraph 2.</p>
			<p>Main paragraph 3.</p>
			<p>Main paragraph 4.</p>
		</main>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := "Main paragraph 1."

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}
