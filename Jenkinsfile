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
                projects = getProjects(rawResults)
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

Project[] getProjects(String rawResults){
    def results = []
    def tempProjects = []
    tempProjects = rawResults.split("\n")
    for(String project in tempProjects){
        results.add(new Project(project.trim())
    }
}