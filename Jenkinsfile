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

            docker.withRegistry('http://150177431474.dkr.ecr.ap-southeast-1.amazonaws.com/products-development', 'ecr:ap-southeast-1:aws-login') {
                app.push("${env.BUILD_NUMBER}")
                app.push("latest")
            }

        }

        stage('Deploy To kubernetes') {


           sh("kubectl --kubeconfig=/var/lib/jenkins/.kube/development-config set image deployment/products-deployment products=150177431474.dkr.ecr.ap-southeast-1.amazonaws.com/products-development:${env.BUILD_NUMBER} --namespace=staging")

           sh("kubectl --kubeconfig=/var/lib/jenkins/.kube/development-config rollout status deployment/products-deployment --namespace=staging")

           sh("kubectl --kubeconfig=/var/lib/jenkins/.kube/development-config rollout history deployment/products-deployment --namespace=staging")

           //Delete the exited container
          // sh("docker rm -v ${docker ps -a -f status=exited -f status=created -q}")

           //Remove the images with none tag
           //sh("docker rmi ${docker images | grep '^<none>' | awk '{ print $3 }'}")

           //removing Dangling images
           //sh(" docker rmi ${docker images -f 'dangling=true' -q}")




        }

    } catch(error) {
        currentBuild.result = "FAILED"
        throw error
    } finally {
         notifyBuild(currentBuild.result)
    }

}

def notifyBuild(String buildStatus = 'STARTED') {
  // build status of null means successful
  buildStatus =  buildStatus ?: 'SUCCESSFUL'

  // Default values
  def colorName = 'RED'
  def colorCode = '#FF0000'
  def subject = "${buildStatus}: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'"
  def summary = "${subject} (${env.BUILD_URL})"

  // Override default values based on build status
  if (buildStatus == 'STARTED') {
    color = 'YELLOW'
    colorCode = '#FFFF00'
  } else if (buildStatus == 'SUCCESSFUL') {
    color = 'GREEN'
    colorCode = '#00FF00'
  } else {
    color = 'RED'
    colorCode = '#FF0000'
  }

  // Send notifications
  slackSend (color: colorCode, message: summary)
}
