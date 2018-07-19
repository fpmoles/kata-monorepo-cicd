#!/usr/bin/env groovy

def String[] projects

try{

    node{

        stage('Checkout'){
            def scmVars = checkout scm
        }
        stage('Collect Details'){
            projects = getProjects()
        }
        stage('Test Artifacts'){

        }
        stage('Build Artifacts'){

        }
        stage('Deploy Artifacts'){

        }
    }
}catch (exc){

}finally{

}


Project[] getProjects(){
        sh "ls -l"
        rawResults=sh(returnStdout: true, script: "ls -l | egrep \'^d\' | awk \'{print \$9}\'")
        echo "Raw Results: ${rawResults}"
        List results=rawResults.split("\n")
        for(String result in results{
            echo "Results value for : ${result}"
        }
        return results;
}

class Project{
    String name
    boolean needsBuild
    boolean testsPass

    Project(name){
        this.name = name
    }
}
