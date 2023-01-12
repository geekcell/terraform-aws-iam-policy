module "codedeploy_policy" {
  source = "../../"

  name = "web-app"
  templates = [
    {
      name = "codedeploy/ecs-blue-green-deployment"
      vars = {
        codedeploy_application_name      = "my-project"
        codedeploy_deployment_group_name = "web-app"

        ecs_cluster_name = "my-project"
        ecs_service_name = "web-app"

        task_definition_task_role_name           = "web-app"
        task_definition_task_execution_role_name = "web-app-exec"
      }
    },
    {
      name = "ecr/push-and-pull"
      vars = {
        ecr_repository = "web-app"
      }
    }
  ]
}
