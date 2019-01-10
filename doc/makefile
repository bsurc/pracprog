all:	pdf html

pdf:
	pandoc -f markdown -t latex syllabus.md -o syllabus.pdf
	pandoc -f markdown -t latex course_desc.md -o desc.pdf
	pandoc -f markdown -t latex design.md -o design.pdf

html:
	markdown syllabus.md > syllabus.html
	markdown course_desc.md > course_desc.html
	markdown design.md > design.html

clean:
	git clean -f

