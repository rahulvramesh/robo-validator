#!groovy
 
/**
        * Author : rahulvramesh@hotmail.com
 */


node {

     sh 'printenv'
 
    environment {
        DISABLE_AUTH = 'true'
        DB_ENGINE    = 'sqlite'
    }

 
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
    }catch(error) {
        //currentBuild.result = "FAILED"
        throw error
    } finally {
        // notifyBuild(currentBuild.result)
    }
}
