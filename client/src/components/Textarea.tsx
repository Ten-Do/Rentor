import { type TextareaHTMLAttributes, forwardRef } from 'react'

type Size = 'sm' | 'md' | 'lg'
type TextareaVariant = 'default' | 'error'
type ColorScheme = 'primary' | 'neutral'

interface TextareaProps
    extends Omit<TextareaHTMLAttributes<HTMLTextAreaElement>, 'size'> {
    size?: Size
    variant?: TextareaVariant
    colorScheme?: ColorScheme
    label?: string
    error?: string | null
}

const sizeClasses: Record<Size, string> = {
    sm: 'px-3 py-2 text-sm rounded-md',
    md: 'px-4 py-3 text-base rounded-lg',
    lg: 'px-5 py-4 text-lg rounded-lg',
}

const baseClasses =
    'w-full bg-gray-800 text-white border transition-colors duration-200 focus:outline-none disabled:opacity-50 disabled:cursor-not-allowed'

const variantClasses: Record<TextareaVariant, Record<ColorScheme, string>> = {
    default: {
        primary: `${baseClasses} border-gray-600 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20`,
        neutral: `${baseClasses} border-gray-600 focus:border-gray-500 focus:ring-2 focus:ring-gray-500/20`,
    },
    error: {
        primary: `${baseClasses} border-red-500 focus:border-red-400 focus:ring-2 focus:ring-red-500/20`,
        neutral: `${baseClasses} border-red-500 focus:border-red-400 focus:ring-2 focus:ring-red-500/20`,
    },
}

export const Textarea = forwardRef<HTMLTextAreaElement, TextareaProps>(
    (
        {
            size = 'md',
            variant = 'default',
            colorScheme = 'primary',
            className = '',
            label,
            id,
            error,
            rows = 5,
            ...props
        },
        ref
    ) => {
        const effectiveVariant = error ? 'error' : variant
        const classes = `${sizeClasses[size]} ${variantClasses[effectiveVariant][colorScheme]} ${className}`

        return (
            <div className='relative group'>
                {label && (
                    <label
                        htmlFor={id}
                        className="block text-sm font-medium text-gray-300 mb-2"
                    >
                        {label}
                    </label>
                )}
                <textarea ref={ref} id={id} className={classes} rows={rows} {...props} />
                {error && (
                    <p className="absolute top-full mt-2 right-0 text-sm text-red-400">{error}</p>
                )}
            </div>
        )
    }
)

Textarea.displayName = 'Textarea'


