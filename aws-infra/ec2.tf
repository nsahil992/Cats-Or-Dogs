resource "aws_instance" "app_server" {
  ami                         = "ami-0f58b397bc5c1f2e8"
  instance_type               = "t2.micro"
  subnet_id                   = aws_subnet.private_subnet.id
  vpc_security_group_ids      = [aws_security_group.ec2_sg.id]
  associate_public_ip_address = false   # private subnet

  key_name = "cod-key"

  tags = {
    Name = "cod-ec2"
  }
}
