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
                tempProjects = rawResults.split("\n")
                for(String project in tempProjects){
                    projects.add(new Project(project.trim())
                }
                echo rawResults
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

class Project{
    String name


    Project(String name){
        this.name = name
        println(this)
    }

    String toString(){
        String result = "[Project: {"
        result = result + "name=" + this.name
        result = result + "}]"
        return result
    }

}