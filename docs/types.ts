/**
 * Represents the possible statuses of a machine learning model.
 */
export type ModelStatus = 'pending' | 'training' | 'trained' | 'failed';

/**
 * Represents the configuration for a machine learning model.
 */
export interface ModelConfig {
  id: string;
  name: string;
  framework: 'tensorflow' | 'pytorch' | 'scikit-learn';
  version: string;
  hyperparameters: Record<string, any>;
  status: ModelStatus;
  createdAt: Date;
  updatedAt: Date;
}

/**
 * Represents the metadata for a trained model.
 */
export interface ModelMetadata {
  accuracy?: number;
  loss?: number;
  precision?: number;
  recall?: number;
  f1Score?: number;
  trainingTime?: number;
  datasetSize?: number;
}

/**
 * Represents a machine learning model with its configuration and metadata.
 */
export interface MachineLearningModel {
  config: ModelConfig;
  metadata?: ModelMetadata;
}

/**
 * Represents the possible events emitted during model training.
 */
export type TrainingEvent = {
  type: 'epochStart' | 'epochEnd' | 'batchStart' | 'batchEnd' | 'trainingComplete';
  timestamp: Date;
  data?: Record<string, any>;
};

/**
 * Represents the response from a model prediction.
 */
export interface PredictionResult {
  input: any;
  output: any;
  confidence?: number;
  modelId: string;
  timestamp: Date;
}

/**
 * Represents the possible error types in model operations.
 */
export type ModelError = {
  code: string;
  message: string;
  details?: any;
};