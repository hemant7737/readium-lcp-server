export class User {
    id:   number;
    alias:    string;
    email:    string;
    password: string | undefined | null; // Hashed password
    clearPassword: string | undefined | null;
}
