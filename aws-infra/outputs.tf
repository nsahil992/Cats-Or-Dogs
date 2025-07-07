output "vpc_id" {
  value = aws_vpc.main.id
}

output "public_subnet_id" {
  value = aws_subnet.public_subnet.id
}

output "ec2_instance_id" {
  value = aws_instance.app_server.id
}

output "alb_sg_id" {
  value = aws_security_group.alb_sg.id
}
