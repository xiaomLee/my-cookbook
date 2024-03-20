.PHONY: deploy

deploy: book
	@echo "====> deploying to github"
	# Delete the ref to avoid keeping history.
	git update-ref -d refs/heads/gh-pages
	rm -rf /tmp/book
	git worktree add -f /tmp/book gh-pages 
	mdbook build
	rm -rf /tmp/book/*
	cp -rp book/* /tmp/book/
	cd /tmp/book && \
    	git add -A && \
    	git commit -m "deployed on $(shell date) by ${USER}" && \
    	git push origin gh-pages