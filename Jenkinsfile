def projectList = []
projectList.add(new ProjectDescriptor("alpha", ProjectType.GOLANG, BuildType.DOCKER))
projectList.add(new ProjectDescriptor("beta", ProjectType.GOLANG, BuildType.DOCKER))
projectList.add(new ProjectDescriptor("gamma", ProjectType.GOLANG, BuildType.DOCKER))

def executeTests(ProjectDescriptor[] projects){
    projects.each{ project->
        switch(project.projectType){
            case GOLANG:
                println "GOLANG: " + project.name
                break
            case JAVASCRIPT:
                println "JAVASCRIPT: " + project.name
                break
            case JAVA:
                println "JAVA: " + project.name
                break
            case STATIC:
                println "STATIC: " + project.name
                break
        }
    }
}

def executeBuild(ProjectDescriptor[] projects){
    projects.each{ project->
        switch(project.buildType){
            case SHELL:
                println "SHELL: " + project.name
                break
            case DOCKER:
                println "DOCKER: " + project.name
                break
        }
    }
}

pipeline{
    agent any
    stages{
        stage ('Test'){
            steps{
                script{
                    executeTests(projectList)
                }
            }
        }
        stage('Build'){
            steps{
                script{
                    executeBuild(projectList)
                }
            }
        }
    }
}

class ProjectDescriptor{
    String name
    ProjectType projectType
    BuildType buildType
    boolean passedTests
    boolean built

    ProjectDescriptor(name, projectType, buildType){
        this.name = name
        this.projectType = projectType
        this.buildType = buildType
    }

}

enum ProjectType{
    GOLANG, JAVASCRIPT, JAVA, STATIC
}

enum BuildType{
    SHELL, DOCKER
}