package main

import (
	"context"
	"fmt"
	"path/filepath"

	packagingv1alpha1 "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/packaging/v1alpha1"
	datapackagingv1alpha1 "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apiserver/apis/datapackaging/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func main() {
	// Create a client according to the canonical example
	// See: https://github.com/kubernetes/client-go/tree/v0.29.0/examples/out-of-cluster-client-configuration

	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		panic(err)
	}

	// Register kapp-controller APIs with the scheme
	scheme := runtime.NewScheme()
	utilruntime.Must(packagingv1alpha1.AddToScheme(scheme))
	utilruntime.Must(datapackagingv1alpha1.AddToScheme(scheme))

	client, err := client.New(config, client.Options{Scheme: scheme})
	if err != nil {
		panic(err)
	}

	// List package installs
	var pkgis packagingv1alpha1.PackageInstallList
	utilruntime.Must(client.List(context.Background(), &pkgis))

	fmt.Print("packageinstalls:\n")
	for _, pkgi := range pkgis.Items {
		fmt.Printf("- %s\n", pkgi.Name)
	}

	// List packages
	var pkgs datapackagingv1alpha1.PackageList
	utilruntime.Must(client.List(context.Background(), &pkgs))

	fmt.Print("packages:\n")
	for _, pkg := range pkgs.Items {
		fmt.Printf("- %s\n", pkg.Name)
	}
}
