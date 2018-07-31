#!groovy

node('master') {

    def app


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


           sh("kubectl --kubeconfig=/var/lib/jenkins/.kube/development-config set image deployment/products-deployment products=150177431474.dkr.ecr.ap-southeast-1.amazonaws.com/products-development:${env.BUILD_NUMBER} --namespace=staging")

           sh("kubectl --kubeconfig=/var/lib/jenkins/.kube/development-config rollout status deployment/products-deployment --namespace=staging")




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
