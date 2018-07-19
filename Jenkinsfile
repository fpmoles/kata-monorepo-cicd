#!/usr/bin/env groovy

def String[] projects

try{
    stage('Collect Details'){
        projects = getProjects()
    }
    stage('Test Artifacts'){

    }
    stage('Build Artifacts'){

    }
    stage('Deploy Artifacts'){

    }
}catch (exc){

}finally{

}


Project[] getProjects(){
    node{
        rawResults=sh(returnStdout: true, script: "ls -l | egrep \'^d\' | awk \'{print $9}\'")
        echo rawResults
        results[]=rawResults.split("\n").trim()
        return results;
    }
}

class Project{
    String name
    boolean needsBuild
    boolean testsPass

    Project(name){
        this.name = name
    }
}
