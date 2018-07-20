pipeline{
    agent any
    stages{
        stage('Prepare'){
            steps{
                script{
                    def rawResults=sh(returnStdout: true, script: "ls -l | egrep \'^d\' | awk \'{print \$9}\'")
                    echo rawResults
                    projects = getProjects(rawResults)
                    printProjects("Prepare")
                }
            }
        }
        stage('Build'){
            steps{
                script{
                    printProjects("Build")
                }
            }
        }
    }
}

def projects = []

class Project{
    String name


    Project(String name){
        this.name = name
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
    for(int i=0; i<tempProjects.size();i++){
        results.add(tempProjects[i].trim())
    }
    return results
}

void printProjects(String branch){
    println "In branch: " + branch
    for(int i=0;i<projects.size();i++){
        println(projects[i])
    }
}


