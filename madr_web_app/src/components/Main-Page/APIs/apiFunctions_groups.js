export const followersNotJoined = async (user_id, group_id, token) => {
    try {
        const response = await fetch('http://localhost:8080/api/groups/followers_not_joined', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${token}`,
            },
            body: JSON.stringify({
                group_id: group_id,
                creator_id: user_id,
            }),
        });

        if (!response.ok) {
            throw new Error('Failed to share deck');
        }

        return response.json()
    } catch (error) {
        throw new Error('Error sharing deck');
    }
}

export const sendInvite = async (creator_id, invitee_id, group_id, token) => {
    try {
        const response = await fetch('http://localhost:8080/api/invite/send', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${token}`,
            },
            body: JSON.stringify({
                group_id: group_id,
                invitee_id: invitee_id,
                creator_id: creator_id,
            }),
        });

        if (!response.ok) {
            throw new Error('Failed to share deck');
        }
    } catch (error) {
        throw new Error('Error sharing deck');
    }
}