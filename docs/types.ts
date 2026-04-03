export interface User {
  id: number;
  name: string;
  email: string;
  password: string;
  role: string;
  createdAt: Date;
  updatedAt: Date;
}

export interface Token {
  id: number;
  userId: number;
  token: string;
  expiresAt: Date;
}

export interface AuthResponse {
  token: string;
  user: User;
  error: string;
}