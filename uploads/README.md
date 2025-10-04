# Uploads Directory

This directory stores all uploaded files from the application.

## Structure

- `images/` - Contains uploaded images (jpg, jpeg, png, gif, webp)
- `files/` - Contains other uploaded files

## File Naming Convention

Files are automatically renamed on upload with the following pattern:
```
YYYYMMDDHHMMSS_<random-id>.<ext>
```

Example: `20251004173045_a8b3c5d1.jpg`

## Important Notes

- This directory should be added to `.gitignore` to avoid committing uploaded files
- Ensure proper permissions are set for file uploads
- Regular backups of this directory are recommended
