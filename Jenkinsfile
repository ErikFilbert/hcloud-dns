node {
    checkout scm

    stage("Run tests") {
        try {
            sh "/usr/local/go/bin/go test -coverprofile cover.txt ."
            archiveArtifacts 'cover.txt'
        } catch (Exception e) {
            error("Failed to build image")
        }
    }
}
