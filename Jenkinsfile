#!groovy
 
/**
        * Author : rahulvramesh@hotmail.com
 */


pipeline {
    agent any

    environment {
        DISABLE_AUTH = 'true'
        DB_ENGINE    = 'sqlite'
    }
 
   

    nodes {
        def app
     
        stage('Environment') {
            steps {
                sh 'printenv'
            }
        }
        stage('Checkout SCM') {
             steps {
                 checkout scm
             }
        }
       stage('Build Image') {
             steps {
                  app = docker.build("products-development")
             }
       }
     
    }
}
