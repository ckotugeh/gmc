export interface User {
  id: number
  full_name: string
  email: string
  role: string
  verification_status: string
  created_at: string
}

export interface AuthResponse {
  message: string
  token: string
}

export interface Community {
  id: number
  name: string
  description: string
  created_at: string
}

export interface Post {
  id: number
  author_id: number
  community_id: number
  title: string
  content: string
  image_url: string
  created_at: string
}

export interface Comment {
  id: number
  post_id: number
  author_id: number
  content: string
  created_at: string
}
