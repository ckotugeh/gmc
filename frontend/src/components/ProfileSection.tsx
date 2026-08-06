import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import api from '../lib/api'
import { useAuth } from '../context/AuthContext'
import type { User } from '../types'

export default function ProfileSection() {
  const { logout } = useAuth()
  const navigate = useNavigate()

  const [profile, setProfile] = useState<User | null>(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  const [editing, setEditing] = useState(false)
  const [editName, setEditName] = useState('')
  const [updateError, setUpdateError] = useState<string | null>(null)
  const [updating, setUpdating] = useState(false)

  useEffect(() => {
    api
      .get<User>('/profile')
      .then((res) => {
        setProfile(res.data)
      })
      .catch((err) => {
        if (err.response?.status === 401) {
          logout()
          navigate('/login')
        } else {
          setError('Could not load profile.')
        }
      })
      .finally(() => setLoading(false))
  }, [logout, navigate])

  const handleEditClick = () => {
    setEditName(profile?.full_name ?? '')
    setUpdateError(null)
    setEditing(true)
  }

  const handleCancelEdit = () => {
    setEditing(false)
    setUpdateError(null)
  }

  const handleSubmitEdit = async (e: React.FormEvent) => {
    e.preventDefault()
    const trimmed = editName.trim()
    if (!trimmed) {
      setUpdateError('Name cannot be empty.')
      return
    }
    setUpdating(true)
    setUpdateError(null)
    try {
      const res = await api.put<User>('/profile', { full_name: trimmed })
      setProfile(res.data)
      setEditing(false)
    } catch (err: unknown) {
      const status = (err as { response?: { status?: number } }).response?.status
      if (status === 401) {
        logout()
        navigate('/login')
      } else {
        setUpdateError('Could not update profile. Please try again.')
      }
    } finally {
      setUpdating(false)
    }
  }

  if (loading) {
    return (
      <div className="bg-white rounded-xl border border-gray-200 p-5">
        <div className="animate-pulse space-y-3">
          <div className="h-4 bg-gray-200 rounded w-1/2" />
          <div className="h-4 bg-gray-200 rounded w-2/3" />
          <div className="h-4 bg-gray-200 rounded w-1/3" />
        </div>
      </div>
    )
  }

  if (error) {
    return (
      <div className="bg-white rounded-xl border border-gray-200 p-5">
        <p className="text-sm text-red-500">{error}</p>
      </div>
    )
  }

  return (
    <div className="bg-white rounded-xl border border-gray-200 p-5">
      <div className="flex items-center justify-between mb-4">
        <h2 className="text-base font-semibold text-gray-900">My Profile</h2>
        {!editing && (
          <button
            onClick={handleEditClick}
            className="text-sm text-blue-600 hover:text-blue-800 transition-colors"
          >
            Edit
          </button>
        )}
      </div>

      {editing ? (
        <form onSubmit={handleSubmitEdit} className="space-y-3">
          <div>
            <label className="block text-xs text-gray-500 mb-1">Full Name</label>
            <input
              type="text"
              value={editName}
              onChange={(e) => setEditName(e.target.value)}
              className="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
              maxLength={255}
            />
          </div>
          {updateError && <p className="text-xs text-red-500">{updateError}</p>}
          <div className="flex gap-2">
            <button
              type="submit"
              disabled={updating}
              className="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 disabled:opacity-50 transition-colors"
            >
              {updating ? 'Saving…' : 'Save'}
            </button>
            <button
              type="button"
              onClick={handleCancelEdit}
              className="px-4 py-2 text-sm text-gray-600 hover:text-gray-900 transition-colors"
            >
              Cancel
            </button>
          </div>
        </form>
      ) : (
        <dl className="space-y-2 text-sm">
          <div>
            <dt className="text-xs text-gray-500">Name</dt>
            <dd className="text-gray-900 font-medium">{profile?.full_name}</dd>
          </div>
          <div>
            <dt className="text-xs text-gray-500">Email</dt>
            <dd className="text-gray-900">{profile?.email}</dd>
          </div>
          <div>
            <dt className="text-xs text-gray-500">Role</dt>
            <dd className="text-gray-900 capitalize">{profile?.role}</dd>
          </div>
          <div>
            <dt className="text-xs text-gray-500">Verification</dt>
            <dd className="text-gray-900 capitalize">{profile?.verification_status}</dd>
          </div>
        </dl>
      )}
    </div>
  )
}
