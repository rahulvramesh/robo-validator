#!groovy
 
/**
        * Author : rahulvramesh@hotmail.com
 */


node {
    def app
    try {
        stage('Checkout SCM') {
            checkout scm
        }
        stage('Build Image') {
            app = docker.build("products-development")
        }
        stage('Test Image') {
           sh "echo 'WE ARE Testing'"
           sh 'printenv'
        }
        stage('Push Image') {
            //Develop
            //Staging
            //Production
            
            switch(GIT_BRANCH) {
               case "master":
                   sh "echo 'master'"
                   break
            }
            
        }
    }
}
