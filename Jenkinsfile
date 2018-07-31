#!groovy

node{

    def app
    
    sh 'printenv'

    try {

        notifyBuild('STARTED')

        stage('Checkout SCM') {
            sh "echo 'Pulling...' + ${env.BRANCH_NAME}"
            checkout scm
        }

        stage('Build Image') {
            app = docker.build("products-development")
        }

        stage('Test Image') {
          //  sh "./vendor/bin/phpunit"
           sh "echo 'WE ARE Testing'"
        }

        stage('Push Image') {

         
        }

        stage('Deploy To kubernetes') {



        }

    } catch(error) {
        currentBuild.result = "FAILED"
        throw error
    } finally {
         notifyBuild(currentBuild.result)
    }

}

def notifyBuild(String buildStatus = 'STARTED') {
  
}
