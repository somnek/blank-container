### Feature:
------------------
# style output:
- using lipgloss or some shit idk

Cmd:
  * remove (remove container(s))
  * tui (tui mode)
  * list (list all blank containers)


# make it binary
> allow multiple flags

# ref
- cli:
  * [charmbracelet/skate](https://github.com/charmbracelet/skate)
  * [trivy](https://github.dev/aquasecurity/trivy)

- docker:
  * lazy-docker
  * go-docker-client

# TODO LIST
------------------
[ ] count flag (blank --count=2)
[-] stop all `container` gracefully before remove them, then proceed to remove `images`
[ ] ignore stopped containers for now (no need to check if is running), irrelavent