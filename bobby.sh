#!/usr/bin/env bash

set -e

FEATURE_ROOT="client/feature"
OUTPUT_ROOT="public/javascript"

echo "🔍 Scanning for TypeScript files..."

find "$FEATURE_ROOT" -type f -name "*.ts" | while read -r file; do
  # Example: client/feature/expense/ExpenseForm/expense_form.ts

  # Strip prefix
  rel_path="${file#${FEATURE_ROOT}/}"
  # expense/ExpenseForm/expense_form.ts

  # Get feature name (first folder)
  feature_name="$(echo "$rel_path" | cut -d'/' -f1)"

  # Get filename
  filename="$(basename "$file" .ts)"

  # Build output path
  out_dir="$OUTPUT_ROOT/$feature_name"
  out_file="$out_dir/$filename.js"

  echo "📦 $file"
  echo "   → Feature: $feature_name"
  echo "   → Output:  $out_file"
  echo
  mkdir -p "$out_dir"

    npx esbuild "$file" \
    --bundle \
    --outfile="$out_file" \
    --minify

done