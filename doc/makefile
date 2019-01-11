all:	pdf html

pdf:
	git rev-parse HEAD >> syllabus.md && pandoc -f markdown -t latex syllabus.md -o syllabus.pdf
	head -n -1 syllabus.md > tmp.md && mv tmp.md syllabus.md
	pandoc -f markdown -t latex course_desc.md -o desc.pdf
	pandoc -f markdown -t latex design.md -o design.pdf

html:
	markdown syllabus.md > syllabus.html
	markdown course_desc.md > course_desc.html
	markdown design.md > design.html

clean:
	git clean -f -x

