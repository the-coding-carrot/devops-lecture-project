terraform {
  required_providers {
    azurerm = {
        source = "hashicorp/azurerm"
        version = "~> 3.0"
    }
  }
}

provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "rg" {
  name = var.resource_group_name
  location = var.location
}

resource "azurerm_kubernetes_cluster" "aks" {
  name = "aks-cluster"
  location = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
  dns_prefix = "aksproject"

  default_node_pool {
    name = "default"
    node_count = 1
    vm_size = "Standard_B2ps_v2"
  }

  identity {
    type = "SystemAssigned"
  }
}