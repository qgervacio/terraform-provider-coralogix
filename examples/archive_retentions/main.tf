terraform {
  required_providers {
    coralogix = {
      version = "~> 1.10"
      source  = "coralogix/coralogix"
    }
  }
}

provider "coralogix" {
  #api_key = "<add your api key here or add env variable CORALOGIX_API_KEY>"
  #env = "<add the environment you want to work at or add env variable CORALOGIX_ENV>"
}

resource "coralogix_archive_retentions" "example" {
  retentions = [
    {
    },
    {
      name = "name_2"
    },
    {
      name = "name_3"
    },
    {
      name = "name_4"
    },
  ]
}

data "coralogix_archive_retentions" "example" {
}