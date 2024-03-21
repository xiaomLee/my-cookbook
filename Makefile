.PHONY: install clean book serve deploy 

book: clean
	mdbook build

serve:
	mdbook serve

deploy: book
	@echo "====> deploying to github"
	# Delete the ref to avoid keeping history.
	git update-ref -d refs/heads/gh-pages
	rm -rf /private/tmp/book
	git worktree add -f /tmp/book gh-pages 
	mdbook build
	rm -rf /tmp/book/*
	cp -rp book/* /tmp/book/
	cd /tmp/book && \
    	git add -A && \
    	git commit -m "deployed on $(shell date) by ${USER}" && \
    	git push origin gh-pages

clean: 
	rm -rf book
	rm -rf /private/tmp/book

install:
	brew install mdbook