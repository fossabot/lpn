package liferay

import internal "github.com/mdelapenya/lpn/internal"

// DXP implementation for Liferay DXP official images
type DXP struct {
	Tag string
}

// GetContainerName returns the name of the container generated by this type of image
func (d DXP) GetContainerName() string {
	return internal.LpnConfig.GetPortalContainerName("dxp")
}

// GetDeployFolder returns the deploy folder under Liferay Home
func (d DXP) GetDeployFolder() string {
	return d.GetLiferayHome() + "/deploy"
}

// GetDockerHubTagsURL returns the URL of the available tags on Docker Hub
func (d DXP) GetDockerHubTagsURL() string {
	return "liferay/dxp"
}

// GetFullyQualifiedName returns the fully qualified name of the image
func (d DXP) GetFullyQualifiedName() string {
	return "docker.io/" + d.GetRepository() + ":" + d.GetTag()
}

// GetLiferayHome returns the Liferay home for DXP
func (d DXP) GetLiferayHome() string {
	return "/opt/liferay"
}

// GetRepository returns the repository for DXP
func (d DXP) GetRepository() string {
	return internal.LpnConfig.GetPortalImageName("dxp")
}

// GetTag returns the tag of the image
func (d DXP) GetTag() string {
	return d.Tag
}

// GetType returns the type of the image
func (d DXP) GetType() string {
	return "dxp"
}

// GetUser returns the user running the main application
func (d DXP) GetUser() string {
	return "liferay"
}
