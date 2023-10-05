terraform {
  required_providers {
    coralogix = {
      version = "~> 1.8"
      source  = "coralogix/coralogix"
    }
  }
}

provider "coralogix" {
  #api_key = "<add your api key here or add env variable CORALOGIX_API_KEY>"
  #env = "<add the environment you want to work at or add env variable CORALOGIX_ENV>"
}

resource "coralogix_hosted_dashboard" dashboard {
  grafana {
    config_json = file("./grafana_dashboard.json")
    folder = coralogix_grafana_folder.test_folder.id
  }
}

resource "coralogix_grafana_folder" "test_folder" {
  title = "Terraform Test Folder"
}

data "coralogix_hosted_dashboard" imported_dashboard {
  uid = coralogix_hosted_dashboard.dashboard.id
}
