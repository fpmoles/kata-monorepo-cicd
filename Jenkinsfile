def projects = []
pipeline{
    agent any
    stages{
        stage('Prepare'){
            steps{
                script{
                    def rawResults=sh(returnStdout: true, script: "ls -l | egrep \'^d\' | awk \'{print \$9}\'")
                    println rawResults
                    projects = getProjects(rawResults)
                    printProjects(projects, "Prepare")
                }
            }
        }
        stage('Test'){
            steps{
                script{
                    parallel(executeTests(projects))
                    printProjects(projects, "Test")
                }
            }
        }
        stage('Build'){
            steps{
                script{
                    printProjects(projects, "Build")
                }
            }
        }
    }
}


class Project{
    String name
    boolean testStatus


    Project(String name){
        this.name = name
    }

    @Override
    public String toString(){
        String result = "[Project: {"
        result = result + "name=" + this.name
        result = result + " testStatus=" + this.testStatus
        result = result + "}]"
        return result
    }

}

Project[] getProjects(String rawResults){
    def results = []
    def tempProjects = []
    tempProjects = rawResults.split("\n")
    for(String project in tempProjects){
        results.add(new Project(project.trim()))
    }
    return results
}

void printProjects(List projects, String branch){
    println "In branch: " + branch
    for(Project project in projects){
        println project.toString()
    }
}

Map executeTests(List projects){
    def projectMap=[:]
    for(Project project in projects){
        projectMap.put(project.name, executeTest(project))
    }
}

void executeTest(Project project){
    boolean result = sh "cd ${project.} && bin/test.sh"
    project.testStatus = result
}


