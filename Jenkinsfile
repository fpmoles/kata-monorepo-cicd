pipeline{
    agent any
    environment{
        def projects = []
    }
    stages{
        stage('Prepare'){
            steps{
                scm checkout
                rawResults=sh(returnStdout: true, script: "ls -l | egrep \'^d\' | awk \'{print \$9}\'")
                echo rawResults
                //projects = getProjects(rawResults)
            }
        }
        stage('Build'){
            steps{
                step{
                    echo "Here"
                }
            }
        }
    }
}

