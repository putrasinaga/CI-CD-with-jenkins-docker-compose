pipeline{
    agent any

    stages{
        stage ("test"){
            steps{
                echo(
                    "testing"
                )
            }
        }
    }

    post{
        always{
            echo("proses sedang berjalan")
        }
        success{
            echo("proses berhasil")
        }
        failure{
            echo("proses gagal")
        }
       cleanup{
            echo("proses selesai")
        }
    }
}