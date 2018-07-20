pipeline{
    agent any
    environment{
        def rawResults = []
    }
    stages{
        stage('Prepare'){
            steps{
                scm checkout
                rawResults=sh(returnStdout: true, script: "ls -l | egrep \'^d\' | awk \'{print \$9}\'")
                echo rawResults
            }
        }
        stage('Build'){
            steps{

            }
        }
    }
}